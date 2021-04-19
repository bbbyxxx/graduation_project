package dals_log

import (
	"lab_device_management_device/dals"
	"lab_device_management_device/models"

	"github.com/jinzhu/gorm"
)

const (
	LogTableName = "multi_log"
)

func SetLog(db *gorm.DB, modelLog *models.MultiLog) (bool, error) {
	var (
		err error
	)

	if db == nil {
		db, _ = dals.GetConn()
	}

	err = db.Table(LogTableName).Create(modelLog).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
