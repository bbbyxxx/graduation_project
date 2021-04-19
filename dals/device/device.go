package dals_device

import (
	"lab_device_management_device/dals"
	"lab_device_management_device/models"
	"lab_device_management_device/service"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	DeviceTableName = "device"
)

func QueryDeviceByDNMIS(conn *gorm.DB, dNMIS []string) (modelDeviceList []*models.Device, err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(DeviceTableName).Where("device_number_model_id in (?) and is_deleted = 0", dNMIS).Find(&modelDeviceList).Error
	if err != nil {
		log.Printf("[QueryDeviceByDNMIS] query db is failed,err:%v", err)
		return nil, err
	}
	return
}

func UpdateDevice(conn *gorm.DB, modelDevice *models.Device, isDeleted int32) (err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	if isDeleted == 1 {
		modelDevice.IsDeleted = 1
	}
	mapArgs := service.ModelDevice2Args(modelDevice)
	err = conn.Table(DeviceTableName).Where("device_number_model_id = ? and is_deleted = 0", modelDevice.DeviceNumberModelId).Update(mapArgs).Error
	if err != nil {
		log.Printf("[UpdateDevice] update is failed,err:%v\n")
		return
	}
	return
}

func AddDevice(conn *gorm.DB, modelDevice *models.Device) (err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return
		}
	}
	err = conn.Table(DeviceTableName).Create(modelDevice).Error
	if err != nil {
		log.Printf("[AddDevice] Create is failed,err:%v\n", err)
		return
	}
	return
}

func QueryDevice(conn *gorm.DB, DeviceNumberModelId string) (res []*models.Device, err error) {
	if conn == nil {
		conn, err = dals.GetConn()
		if err != nil {
			return nil, err
		}
	}
	res = make([]*models.Device, 0)
	err = conn.Table(DeviceTableName).Where("device_number_model_id = ? and is_deleted = 0", DeviceNumberModelId).Find(&res).Error
	if err != nil {
		log.Printf("[QueryDevice] query DeviceNumberModelId is failed,err:%v\n", err)
		return
	}
	log.Println(len(res))
	return
}
