package paper_const

import "errors"

const (
	AddPaperFailed        = "添加论文失败"
	AddPaperSuccessful    = "添加论文成功"
	PaperNotExist         = "论文不存在"
	PaperHasExist         = "论文已存在"
	UpdatePaperFailed     = "更新论文信息失败"
	UpdatePaperSuccessful = "更新论文信息成功"
	QueryPaperFailed      = "查询论文信息失败"
	QueryPaperSuccessful  = "查询论文信息成功"
	DeletePaperFailed     = "删除论文信息失败"
	DeletePaperSuccessful = "删除论文信息成功"
)

const (
	AddNewPaper = "新增论文"
	DeletePaper = "删除论文"
	UpdatePaper = "更新论文"
)

var (
	QueryPaperFailedErr = errors.New("查询论文信息失败")
)
