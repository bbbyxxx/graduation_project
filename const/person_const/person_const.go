package person_const

import "errors"

var (
	PersonNotAllowRepeatRegist = errors.New("不能重复注册")
	LoginError                 = errors.New("登录失败")
)

var (
	PersonRegistSuccessful = "注册成功"
	PersonRegistFailed     = "注册失败"
	PersonLoginSuccessful  = "登录成功"
	PersonLoginFailed      = "登录失败"
	PersonUpdateSuccessful = "更新成功"
	PersonUpdateFailed     = "更新失败"
	PersonDeleteSuccessful = "删除成功"
	PersonDeleteFailed     = "删除失败"
	PersonNotExist         = "此人不存在"
)

var (
	Update  = "更新"
	Add     = " 新增"
	Query   = "查询"
	Deleted = "删除"
)

var (
	Regist = "注册"
)
