package models

import "time"

type Device struct {
	DeviceNumberId      string    `json:"device_number_id"`
	DeviceModelId       string    `json:"device_model_id"`
	DeviceNumberModelId string    `json:"device_number_model_id"`
	DeviceName          string    `json:"device_name"`
	DeviceStatus        int32     `json:"device_status"`
	DeviceImages        string    `json:"device_images"`
	DeviceUseDesc       string    `json:"device_use_desc"`
	DeviceCreateTime    time.Time `json:"device_create_time"`
	LaboratoryId        int32     `json:"laboratory_id"`
	LaboratoryFloor     int32     `json:"laboratory_floor"`
	LaboratoryFloorId   string    `json:"laboratory_floor_id"`
	IsDeleted           int32     `json:"is_deleted"`
}
