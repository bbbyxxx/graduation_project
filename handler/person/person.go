package handler_person

import (
	"context"
	"encoding/json"
	"lab_device_management_person/const/person_const"
	dal_log "lab_device_management_person/dals/log"
	dal_person "lab_device_management_person/dals/person"

	"lab_device_management_person/dals"
	"lab_device_management_person/models"
	"lab_device_management_person/proto/person/person"
	"log"
	"time"
)

type Person struct{}

func (p *Person) RegisterPerson(ctx context.Context, personReq *person.RegisterRequest) (*person.RegisterResponse, error) {
	var (
		res  *person.RegisterResponse
		succ bool
		err  error
	)
	res = &person.RegisterResponse{}
	res.Message = person_const.PersonRegistFailed
	modelPerson := &models.Person{
		MultiId:    personReq.Person.MultiId,
		Name:       personReq.Person.Name,
		Sex:        personReq.Person.Sex,
		Password:   personReq.Person.Password,
		Phone:      personReq.Person.Phone,
		Major:      personReq.Person.Major,
		Grade:      personReq.Person.Grade,
		Class:      personReq.Person.Class,
		RegistTime: time.Unix(time.Now().Unix(), 0),
		UpdateTime: time.Unix(time.Now().Unix(), 0),
		Indentity:  personReq.Person.Indentity,
	}
	modelPersonQuery, err := dal_person.QueryPersonByMultiIdAndIndentity(nil, modelPerson)
	if err != nil {
		log.Printf("[RegisterPerson] QueryPersonByMultiIdAndIndentity is failed,err:%v", err)
		return res, err
	}
	//存在，不能重复注册
	if len(modelPersonQuery) != 0 {
		return res, person_const.PersonNotAllowRepeatRegist
	}
	succ, err = dal_person.AddPerson(nil, modelPerson)
	if !succ {
		log.Printf("[RegisterPerson] AddPerson is failed,err:%v", err)
		return res, err
	}
	//添加日志
	logContent, _ := json.Marshal(modelPerson)
	var modelLog = &models.MultiLog{
		Operator:            person_const.Regist,
		MultiId:             modelPerson.MultiId,
		DeviceNumberModelId: "",
		BeforeContent:       "",
		AfterContent:        string(logContent),
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	succ, err = dal_log.SetLog(nil, modelLog)
	if !succ {
		log.Printf("[RegisterPerson] SetLog is failed,err:%v\n", err)
		return res, err
	}
	res.Message = person_const.PersonRegistSuccessful
	return res, nil
}

//更新删除接口放到一块，根据请求字段is_deleted进行区分
func (p *Person) UpdatePerson(ctx context.Context, personReq *person.UpdateRequest) (*person.UpdateResponse, error) {
	var (
		res       *person.UpdateResponse
		isDeleted bool
		succ      bool
		err       error
		operate   string
	)
	res = &person.UpdateResponse{}
	conn, err := dals.GetConn()
	if err != nil {
		log.Printf("[UpdatePerson] QueryPersonByMultiIdAndIndentity is failed,err:%v\n", err)
		return res, err
	}
	tx := conn.Begin()
	tx = tx.Set("gorm:query_option", "FOR UPDATE")
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	res.Message = person_const.PersonUpdateFailed
	operate = person_const.Update
	modelPerson := &models.Person{
		MultiId:    personReq.Person.MultiId,
		Name:       personReq.Person.Name,
		Sex:        personReq.Person.Sex,
		Password:   personReq.Person.Password,
		Phone:      personReq.Person.Phone,
		Major:      personReq.Person.Major,
		Grade:      personReq.Person.Grade,
		Class:      personReq.Person.Class,
		UpdateTime: time.Unix(time.Now().Unix(), 0),
		Indentity:  personReq.Person.Indentity,
	}
	//更新时，先查询，看一下人物是否存在
	modelPersonOld, err := dal_person.QueryPersonByMultiIdAndIndentity(tx, modelPerson)
	if err != nil {
		log.Printf("[UpdatePerson] QueryPersonByMultiIdAndIndentity is failed,err:%v\n", err)
		res.Message = person_const.PersonNotExist
		return res, err
	}
	if personReq.IsDeleted == 1 {
		isDeleted = true
		operate = person_const.Deleted
	}
	//在进行更新
	succ, err = dal_person.UpdatePerson(tx, modelPerson, isDeleted)
	if !succ {
		log.Printf("[UpdatePerson] UpdatePerson is failed,err:%v\n", err)
		return res, err
	}
	//添加日志
	logContent, _ := json.Marshal(modelPerson)
	logContentOld, _ := json.Marshal(modelPersonOld)
	var modelLog = &models.MultiLog{
		Operator:            operate,
		MultiId:             modelPerson.MultiId,
		DeviceNumberModelId: "",
		BeforeContent:       string(logContentOld),
		AfterContent:        string(logContent),
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	succ, err = dal_log.SetLog(tx, modelLog)
	if !succ {
		log.Printf("[RegisterPerson] SetLog is failed,err:%v\n", err)
		return res, err
	}

	res.Message = person_const.PersonUpdateSuccessful
	return res, nil
}

func (p *Person) Login(ctx context.Context, personReq *person.LoginRequest) (*person.LoginResponse, error) {
	var (
		res *person.LoginResponse
		err error
	)
	res = &person.LoginResponse{}
	res.Message = person_const.PersonLoginFailed
	modelPerson, err := dal_person.LoginValid(nil, personReq.MultiId, personReq.Password)
	if err != nil {
		log.Printf("[Login] LoginValid is failed,err:%v\n", err)
		return res, err
	}
	if len(modelPerson) == 0 {
		log.Println("[Login] len(modelPerson) is 0")
		return res, person_const.LoginError
	}
	if modelPerson[0].MultiId == personReq.MultiId {
		log.Printf("login in %v\n", time.Now())
	}
	modelPerson[0].LastLoginTime = modelPerson[0].LoginTime
	modelPerson[0].LoginTime = time.Unix(time.Now().Unix(), 0)
	_, err = dal_person.UpdateLoginTimePerson(nil, modelPerson[0], false)
	if err != nil {
		log.Printf("[Login] UpdatePerson is failed,err:%v\n", err)
		return res, err
	}
	res.Message = person_const.PersonLoginSuccessful
	return res, nil
}
