package device_const

import "errors"

const (
	AddDeviceFailed     = "添加设备失败"
	AddDeviceSuccessful = "添加设备成功"
	AddDeviceRepeat     = "不能重复添加相同设备"

	DeleteDeviceFailed     = "删除设备失败"
	DeleteDeviceSuccessful = "删除设备成功"
	DeleteDeviceNotExist   = "待删除设备不存在"

	UpdateDeviceFailed     = "更新设备失败"
	UpdateDeviceSuccessful = "更新设备成功"
	UpdateDeviceNotExist   = "待更新设备不存在"

	QueryDeviceFailed     = "查询设备失败"
	QueryDeviceSuccessful = "查询设备成功"
	QueryDeviceNotExist   = "设备不存在"
)

const (
	LogAddDevice    = "添加设备"
	LogDeleteDevice = "删除设备"
	LogUpdateDevice = "更新设备"
)

var (
	AddDeviceFailedErr     = errors.New("添加设备失败")
	AddDeviceSuccessfulErr = errors.New("添加设备成功")
	AddDeviceRepeatErr     = errors.New("不能重复添加相同设备")

	DeleteDeviceFailedErr     = errors.New("删除设备失败")
	DeleteDeviceSuccessfulErr = errors.New("删除设备成功")
	DeleteDeviceNotExistErr   = errors.New("待删除设备不存在")

	UpdateDeviceFailedErr     = errors.New("更新设备失败")
	UpdateDeviceSuccessfulErr = errors.New("更新设备成功")
	UpdateDeviceNotExistErr   = errors.New("待更新设备不存在")

	QueryDeviceFailedErr     = errors.New("查询设备信息失败")
	QueryDeviceSuccessfulErr = errors.New("查询设备信息成功")
	QueryDeviceNotExistErr   = errors.New("设备不存在")
)
