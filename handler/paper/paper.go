package handler_paper

import (
	"context"
	"encoding/json"
	"errors"
	paper_const "lab_device_management_device/const/paper"
	"lab_device_management_device/dals"
	dals_log "lab_device_management_device/dals/log"
	dals_paper "lab_device_management_device/dals/paper"
	"lab_device_management_device/models"
	"lab_device_management_device/proto/paper/paper"
	"log"
	"strings"
	"time"
)

type Paper struct{}

//支持通过paper_number直接查询论文及multi_id和device_number_model_id查询或者multi_id查询与此人相关的论文
func (p *Paper) MGetPaper(ctx context.Context, reqPaper *paper.MGetPaperRequest) (*paper.MGetPaperResponse, error) {
	var (
		resp      *paper.MGetPaperResponse
		paperList []*paper.Paper
		err       error
	)
	resp = &paper.MGetPaperResponse{}
	resp.Message = paper_const.QueryPaperFailed

	modelPersonDevicePaper := &models.PersonDevicePaper{
		MultiId:             reqPaper.MultiId,
		DeviceNumberModelId: reqPaper.DeviceNumberModelId,
		PaperNumber:         reqPaper.PaperNumber,
	}
	modelPaperList, err := dals_paper.MGetPaper(nil, modelPersonDevicePaper)
	if err != nil {
		return resp, nil
	}
	if len(modelPaperList) == 0 {
		return resp, paper_const.QueryPaperFailedErr
	}

	paperList = make([]*paper.Paper, 0)
	for _, modelPaper := range modelPaperList {
		paper := &paper.Paper{
			PaperNumber:         modelPaper.PaperNumber,
			PaperName:           modelPaper.PaperName,
			PaperTopic:          modelPaper.PaperTopic,
			PaperContent:        modelPaper.PaperContent,
			RelatedCode:         modelPaper.RelatedCode,
			DeviceNumberModelId: modelPaper.DeviceNumberModelId,
		}
		paperList = append(paperList, paper)
	}
	resp.Message = paper_const.QueryPaperSuccessful
	resp.Paper = paperList
	return resp, nil
}

func (p *Paper) AddPaper(ctx context.Context, reqPaper *paper.AddPaperRequest) (*paper.AddPaperResponse, error) {
	var (
		resp          *paper.AddPaperResponse
		err           error
		recordIsExist bool //标记 用户-设备-论文表中是否已有请求记录存在
	)
	resp = &paper.AddPaperResponse{}
	resp.Message = paper_const.AddPaperFailed

	err = checkParams(reqPaper)
	if err != nil {
		log.Printf("[AddPaper] checkParams is failed,err:%v\n", err)
		return resp, err
	}

	//1.先查询 paper_number 的论文是否存在及在 用户-设备-论文表中 有没有记录
	modelPaperList, err := dals_paper.GetPaperByPaperNumber(nil, reqPaper.Paper.PaperNumber)
	if err != nil {
		return resp, err
	}
	if len(modelPaperList) > 0 {
		resp.Message = paper_const.PaperHasExist
		return resp, err
	}
	//不根据论文编号进行查询，因为此时论文还没有加入
	args := map[string]interface{}{
		"multi_id":               reqPaper.MultiId,
		"device_number_model_id": reqPaper.Paper.DeviceNumberModelId,
	}
	pdpList, err := dals_paper.GetPersonDevicePaper(nil, args)
	if err != nil {
		return resp, err
	}
	if len(pdpList) > 0 {
		recordIsExist = true
	}

	modelPaper := &models.Paper{
		PaperNumber:         reqPaper.Paper.PaperNumber,
		PaperName:           reqPaper.Paper.PaperName,
		PaperTopic:          reqPaper.Paper.PaperTopic,
		PaperContent:        reqPaper.Paper.PaperContent,
		RelatedCode:         reqPaper.Paper.RelatedCode,
		DeviceNumberModelId: reqPaper.Paper.DeviceNumberModelId,
	}

	db, _ := dals.GetConn()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	//2.插入库中
	err = dals_paper.AddPaper(tx, modelPaper)
	if err != nil {
		return resp, err
	}

	//3.如果在 用户-设备-论文表中 没有记录，则插入全新映射关系；否则,给paper_number字段加上就行了
	modelPersonDevicePaper := &models.PersonDevicePaper{
		MultiId:             reqPaper.MultiId,
		DeviceNumberModelId: modelPaper.DeviceNumberModelId,
		PaperNumber:         modelPaper.PaperNumber,
	}
	err = dals_paper.AddPersonDevicePaper(tx, args, modelPersonDevicePaper, recordIsExist)
	if err != nil {
		return resp, err
	}

	//4.记录日志
	beforeContent, _ := json.Marshal(modelPaperList[0])
	newContent, _ := json.Marshal(modelPaper)
	modelLog := &models.MultiLog{
		Operator:            paper_const.AddNewPaper,
		MultiId:             reqPaper.MultiId,
		DeviceNumberModelId: modelPaper.DeviceNumberModelId,
		BeforeContent:       string(beforeContent),
		AfterContent:        string(newContent),
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	flag, err := dals_log.SetLog(tx, modelLog)
	if err != nil || !flag {
		return resp, err
	}
	resp.Message = paper_const.AddPaperSuccessful
	return resp, nil
}

func (p *Paper) UpdatePaper(ctx context.Context, reqPaper *paper.UpdatePaperRequest) (*paper.UpdatePaperResponse, error) {
	var (
		resp     *paper.UpdatePaperResponse
		err      error
		operator string
	)
	resp = &paper.UpdatePaperResponse{}
	if reqPaper.IsDeleted == 0 {
		resp.Message = paper_const.UpdatePaperFailed
		operator = paper_const.UpdatePaper
	} else {
		resp.Message = paper_const.DeletePaperFailed
		operator = paper_const.DeletePaper
	}
	//1.查询待更新记录是否存在（根据论文编号），存在，更新部分字段；不存在,提示此记录不存在
	modelPaperList, err := dals_paper.GetPaperByPaperNumber(nil, reqPaper.PaperNumber)
	if err != nil {
		return resp, err
	}
	if len(modelPaperList) == 0 {
		resp.Message = paper_const.PaperNotExist
		return resp, err
	}

	db, _ := dals.GetConn()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	//2.进行更新操作
	args := getUpdatePaperArgs(reqPaper)
	err = dals_paper.UpdatePaper(tx, args, reqPaper.PaperNumber)
	if err != nil {
		return resp, nil
	}

	//3.如果是删除操作的话，设置 用户-设备-论文 表中的论文列为空
	if reqPaper.IsDeleted == 1 {
		args := map[string]interface{}{
			"paper_number": "",
		}
		modelPersonDevicePaper := &models.PersonDevicePaper{
			MultiId:             reqPaper.MultiId,
			DeviceNumberModelId: reqPaper.DeviceNumberModelId,
			PaperNumber:         reqPaper.PaperNumber,
		}
		err = dals_paper.UpdatePersonDevicePaper(tx, args, modelPersonDevicePaper)
		if err != nil {
			return resp, err
		}
	}

	//4.记录日志操作,不采用参数值，因为参数值可能为空
	modelPaperListNew, err := dals_paper.GetPaperByPaperNumber(nil, reqPaper.PaperNumber)
	if err != nil {
		return resp, err
	}
	beforeContent, _ := json.Marshal(modelPaperList[0])
	newContent, _ := json.Marshal(modelPaperListNew[0])
	modelLog := &models.MultiLog{
		Operator:            operator,
		MultiId:             reqPaper.MultiId,
		DeviceNumberModelId: reqPaper.DeviceNumberModelId,
		BeforeContent:       string(beforeContent),
		AfterContent:        string(newContent),
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	flag, err := dals_log.SetLog(tx, modelLog)
	if err != nil || !flag {
		return resp, err
	}

	if reqPaper.IsDeleted == 0 {
		resp.Message = paper_const.UpdatePaperSuccessful
	} else {
		resp.Message = paper_const.DeletePaperSuccessful
	}
	return resp, nil
}

func checkParams(reqPaper *paper.AddPaperRequest) error {
	if s := strings.TrimSpace(reqPaper.MultiId); s == "" {
		return errors.New("请求参数multi_id不能为空")
	}

	if s := strings.TrimSpace(reqPaper.Paper.DeviceNumberModelId); s == "" {
		return errors.New("device_number_model_id")
	}

	if s := strings.TrimSpace(reqPaper.Paper.PaperNumber); s == "" {
		return errors.New("论文编号不能为空")
	}
	if s := strings.TrimSpace(reqPaper.Paper.PaperName); s == "" {
		return errors.New("论文名称不能为空")
	}
	if s := strings.TrimSpace(reqPaper.Paper.PaperTopic); s == "" {
		return errors.New("论文主题不能为空")
	}
	if s := strings.TrimSpace(reqPaper.Paper.PaperContent); s == "" {
		return errors.New("论文内容不能为空")
	}
	if s := strings.TrimSpace(reqPaper.Paper.RelatedCode); s == "" {
		return errors.New("论文相关代码不能为空")
	}
	return nil
}

func getUpdatePaperArgs(reqPaper *paper.UpdatePaperRequest) map[string]interface{} {
	args := make(map[string]interface{})
	if s := strings.TrimSpace(reqPaper.DeviceNumberModelId); s != "" {
		args["device_number_model_id"] = s
	}
	if s := strings.TrimSpace(reqPaper.PaperNumber); s != "" {
		args["paper_number"] = s
	}
	if s := strings.TrimSpace(reqPaper.RelatedCode); s != "" {
		args["related_code"] = s
	}
	if s := strings.TrimSpace(reqPaper.PaperContent); s != "" {
		args["paper_content"] = s
	}
	if s := strings.TrimSpace(reqPaper.PaperTopic); s != "" {
		args["paper_topic"] = s
	}
	if s := strings.TrimSpace(reqPaper.PaperName); s != "" {
		args["paper_name"] = s
	}
	if reqPaper.IsDeleted == 1 {
		args["is_deleted"] = 1
	}
	return args
}
