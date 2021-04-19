package dals_paper

import (
	"lab_device_management_device/dals"
	"lab_device_management_device/models"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	PaperTableName             = "paper"
	PersonDevicePaperTableName = "person_device_paper"
)

func MGetPaper(conn *gorm.DB, modelPersonDevicePaper *models.PersonDevicePaper) (modelPaperList []*models.Paper, err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	conn = conn.Table(PersonDevicePaperTableName).Select("paper_number")
	if modelPersonDevicePaper.MultiId != "" && modelPersonDevicePaper.DeviceNumberModelId != "" {
		conn = conn.Where("multi_id = ? and device_number_model_id = ?", modelPersonDevicePaper.MultiId, modelPersonDevicePaper.DeviceNumberModelId)
	} else if modelPersonDevicePaper.PaperNumber != "" {
		conn = conn.Where("paper_number = ?", modelPersonDevicePaper.PaperNumber)
	} else if modelPersonDevicePaper.MultiId != "" {
		conn = conn.Where("multi_id = ?", modelPersonDevicePaper.MultiId)
	} else if modelPersonDevicePaper.DeviceNumberModelId != "" {
		conn = conn.Where("device_number_model_id = ?", modelPersonDevicePaper.DeviceNumberModelId)
	}
	var paperNumberList []string
	err = conn.Find(&paperNumberList).Error
	if err != nil {
		log.Printf("[MGetPaper] query is failed,err:%v\n", err)
		return
	}
	newPaperNumberList := make([]string, 0)
	for _, paperNumber := range paperNumberList {
		if paperNumber == "" {
			continue
		}
		newPaperNumberList = append(newPaperNumberList, paperNumber)
	}
	err = conn.Table(PaperTableName).Where("paper_number in (?)", newPaperNumberList).Find(&modelPaperList).Error
	if err != nil {
		log.Printf("[MGetPaper] quety paper is failed,err:%v\n", err)
		return
	}
	return
}

func AddPersonDevicePaper(conn *gorm.DB, args map[string]interface{}, modelPersonDevicePaper *models.PersonDevicePaper, isExist bool) (err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	if isExist {
		err = conn.Table(PersonDevicePaperTableName).Where(args).Update(modelPersonDevicePaper).Error
		if err != nil {
			log.Printf("[AddPersonDevicePaper] update record is failed,err:%v\n", err)
			return
		}
	} else {
		err = conn.Table(PersonDevicePaperTableName).Create(modelPersonDevicePaper).Error
		if err != nil {
			log.Printf("[AddPersonDevicePaper] create record is failed,err:%v\n", err)
			return
		}
	}
	return
}

func UpdatePersonDevicePaper(conn *gorm.DB, args map[string]interface{}, modelPaper *models.PersonDevicePaper) (err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(PersonDevicePaperTableName).Where(modelPaper).Update(args).Error
	if err != nil {
		log.Printf("[UpdatePersonDevicePaper] update is failed,err:%v\n", err)
		return
	}
	return
}

func GetPersonDevicePaper(conn *gorm.DB, args map[string]interface{}) (list []*models.PersonDevicePaper, err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(PersonDevicePaperTableName).Where(args).Find(&list).Error
	if err != nil {
		log.Printf("[GetPersonDevicePaper] query is failed,err:%v\n", err)
		return
	}
	return
}

func GetPaperByPaperNumber(conn *gorm.DB, paperNumber string) (modelPaperList []*models.Paper, err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(PaperTableName).Where("paper_number = ? and is_deleted = 0", paperNumber).Find(&modelPaperList).Error
	if err != nil {
		log.Printf("[GetPaperByPaperNumber] Query is failed,err:%v\n", err)
		return
	}
	return
}

func AddPaper(conn *gorm.DB, modelPaper *models.Paper) (err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(PaperTableName).Create(modelPaper).Error
	if err != nil {
		log.Printf("[AddPaper] CREATE is failed,err:%v\n", err)
		return
	}
	return
}

func UpdatePaper(conn *gorm.DB, args map[string]interface{}, paperNumber string) (err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(PaperTableName).Where("paper_number = ?", paperNumber).Update(args).Error
	if err != nil {
		log.Printf("[UpdatePaper] update is failed,err:%v\n", err)
		return
	}
	return
}
