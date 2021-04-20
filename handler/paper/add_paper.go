package handler_paper

import (
	"errors"
	"lab_device_management_api/model"
	"lab_device_management_api/rpc"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AddPaper(c *gin.Context) {
	var (
		modelPaper model.Paper
		err        error
	)
	c.ShouldBindJSON(&modelPaper)
	log.Printf("modelDevice is %v\n", modelPaper)
	multiIdInterface, exist := c.Get("multi_id")
	multiId, _ := multiIdInterface.(string)
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "未登录",
			"data":    "",
		})
		return
	}

	//1.检查参数
	err = checkParams(&modelPaper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
	}

	//2.调用rpc服务进行添加论文信息
	resp, err := rpc.AddPaper(c, nil, &modelPaper, multiId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
		"data":    "",
	})
	return
}

func checkParams(modelPaper *model.Paper) error {
	if modelPaper.IsDeleted == 0 {
		if s := strings.TrimSpace(modelPaper.PaperName); s == "" {
			return errors.New("论文名称不能为空")
		}
		if s := strings.TrimSpace(modelPaper.PaperTopic); s == "" {
			return errors.New("论文主题不能为空")
		}
		if s := strings.TrimSpace(modelPaper.PaperContent); s == "" {
			return errors.New("论文内容不能为空")
		}
		if s := strings.TrimSpace(modelPaper.RelatedCode); s == "" {
			return errors.New("论文相关代码不能空")
		}
		if s := strings.TrimSpace(modelPaper.DeviceNumberModelId); s == "" {
			return errors.New("设备编号-型号不能为空")
		}
	}
	if s := strings.TrimSpace(modelPaper.PaperNumber); s == "" {
		return errors.New("论文编号不能为空")
	}
	return nil
}
