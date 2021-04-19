package models

import "time"

type MultiLog struct {
	Operator            string    `json:"operator"`
	MultiId             string    `json:"multi_id"`
	DeviceNumberModelId string    `json:"device_number_model_id"`
	BeforeContent       string    `json:"before_content"`
	AfterContent        string    `json:"after_content"`
	OperateTime         time.Time `json:"operate_time"`
}
