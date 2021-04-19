package models

import "time"

type Person struct {
	MultiId       string    `json:"multi_id"`
	Name          string    `json:"name"`
	Sex           string    `json:"sex"`
	Password      string    `json:"password"`
	Phone         string    `json:"phone"`
	Major         string    `json:"major"`
	Grade         int32     `json:"grade"`
	Class         int32     `json:"class"`
	RegistTime    time.Time `json:"regist_time"`
	UpdateTime    time.Time `json:"update_time"`
	LoginTime     time.Time `json:"login_time"`
	LastLoginTime time.Time `json:"last_login_time"`
	Indentity     int32     `json:"indentity"`
	IsDeleted     int32     `json:"is_deleted"`
}
