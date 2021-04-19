package model

type Person struct {
	MultiId       string `json:"multi_id"`
	Name          string `json:"name"`
	Sex           string `json:"sex"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Major         string `json:"major"`
	Grade         int32  `json:"grade"`
	Class         int32  `json:"class"`
	RegistTime    int64  `json:"regist_time"`
	UpdateTime    int64  `json:"update_time"`
	LoginTime     int64  `json:"login_time"`
	LastLoginTime int64  `json:"last_login_time"`
	Indentity     int32  `json:"indentity"`
	IsDeleted     int32  `json:"is_deleted"`
}

type Login struct {
	MultiId   string `json:"multi_id"`
	Password  string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Code      int32  `json:"code"`
}

type Token struct {
	MultiId string `json:"multi_id"`
	Token   string `json:"token"`
}
