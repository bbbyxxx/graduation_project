package handler_device

import (
	"context"
	"encoding/json"
	"fmt"
	device2 "lab_device_management_device/const/device"
	"lab_device_management_device/dals"
	dals_device "lab_device_management_device/dals/device"
	dals_log "lab_device_management_device/dals/log"
	"lab_device_management_device/models"
	"lab_device_management_device/proto/device/device"
	"log"
	"time"
)

type Device struct{}

func (d *Device) AddDevice(ctx context.Context, reqDevice *device.AddDeviceRequest) (*device.AddDeviceResponse, error) {
	var (
		resp *device.AddDeviceResponse
		err  error
	)
	resp = &device.AddDeviceResponse{}
	resp.Message = device2.AddDeviceFailed

	//1.根据DeviceNumberModelId先去数据库查询，若存在，则直接返回设备已存在
	DeviceNumberModelId := reqDevice.Device.DeviceNumberId + "-" + reqDevice.Device.DeviceModelId
	log.Println(DeviceNumberModelId)
	modelDeviceList, err := dals_device.QueryDevice(nil, DeviceNumberModelId)
	if err != nil {
		return resp, err
	}
	if len(modelDeviceList) > 0 {
		resp.Message = device2.AddDeviceRepeat
		return resp, device2.AddDeviceRepeatErr
	}

	db, _ := dals.GetConn()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	//2.封装数据，插入数据库
	laboratoryFloorId := fmt.Sprintf("%d-%d", reqDevice.Device.LaboratoryId, reqDevice.Device.LaboratoryFloor)
	modelDevice := &models.Device{
		DeviceNumberId:      reqDevice.Device.DeviceNumberId,
		DeviceModelId:       reqDevice.Device.DeviceModelId,
		DeviceNumberModelId: DeviceNumberModelId,
		DeviceName:          reqDevice.Device.DeviceName,
		DeviceStatus:        reqDevice.Device.DeviceStatus, //完好无人使用
		DeviceImages:        reqDevice.Device.DeviceImages,
		DeviceUseDesc:       reqDevice.Device.DeviceUseDesc,
		DeviceCreateTime:    time.Unix(time.Now().Unix(), 0),
		LaboratoryId:        reqDevice.Device.LaboratoryId,
		LaboratoryFloor:     reqDevice.Device.LaboratoryFloor,
		LaboratoryFloorId:   laboratoryFloorId,
	}
	err = dals_device.AddDevice(tx, modelDevice)
	if err != nil {
		return resp, err
	}

	content, _ := json.Marshal(modelDevice)
	//3.记录操作日志
	modelLog := &models.MultiLog{
		Operator:            device2.LogAddDevice,
		MultiId:             reqDevice.MultiId,
		DeviceNumberModelId: DeviceNumberModelId,
		BeforeContent:       "",
		AfterContent:        string(content),
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	flag, err := dals_log.SetLog(tx, modelLog)
	if err != nil || !flag {
		return resp, err
	}

	resp.Message = device2.AddDeviceSuccessful
	return resp, nil
}

func (d *Device) DeleteDevice(ctx context.Context, reqDevice *device.DeleteDeviceRequest) (*device.DeleteDeviceResponse, error) {
	var (
		resp *device.DeleteDeviceResponse
		err  error
	)
	resp = &device.DeleteDeviceResponse{}
	resp.Message = device2.DeleteDeviceFailed
	//1.根据DeviceNumberModelId查询设备，是否存在
	modelDeviceList, err := dals_device.QueryDevice(nil, reqDevice.DeviceNumberModelId)
	if err != nil {
		return resp, err
	}
	if len(modelDeviceList) == 0 {
		resp.Message = device2.DeleteDeviceNotExist
		return resp, device2.DeleteDeviceNotExistErr
	}

	db, err := dals.GetConn()
	if err != nil {
		return resp, err
	}
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	//2.删除设备（软删除）
	err = dals_device.UpdateDevice(tx, modelDeviceList[0], reqDevice.IsDeleted)
	if err != nil {
		return resp, err
	}

	//3.记录日志
	beforeContent, _ := json.Marshal(modelDeviceList[0])
	modelLog := &models.MultiLog{
		Operator:            device2.LogDeleteDevice,
		MultiId:             reqDevice.MultiId,
		DeviceNumberModelId: reqDevice.DeviceNumberModelId,
		BeforeContent:       string(beforeContent),
		AfterContent:        "",
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	flag, err := dals_log.SetLog(tx, modelLog)
	if err != nil || !flag {
		return resp, err
	}

	resp.Message = device2.DeleteDeviceSuccessful
	return resp, nil
}

func (d *Device) UpdateDevice(ctx context.Context, reqDevice *device.UpdateDeviceRequest) (*device.UpdateDeviceResponse, error) {
	var (
		resp *device.UpdateDeviceResponse
		err  error
	)
	resp = &device.UpdateDeviceResponse{}
	resp.Message = device2.UpdateDeviceFailed
	//1.根据DeviceNumberModelId查询设备，是否存在
	deviceNumberModelId := reqDevice.Device.DeviceNumberId + "-" + reqDevice.Device.DeviceModelId
	modelDeviceList, err := dals_device.QueryDevice(nil, deviceNumberModelId)
	if err != nil {
		return resp, err
	}
	if len(modelDeviceList) == 0 {
		resp.Message = device2.UpdateDeviceNotExist
		return resp, device2.UpdateDeviceNotExistErr
	}

	db, err := dals.GetConn()
	if err != nil {
		return resp, err
	}
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	//2.更新设备

	modelDevice := &models.Device{
		DeviceNumberId:      reqDevice.Device.DeviceNumberId,
		DeviceModelId:       reqDevice.Device.DeviceModelId,
		DeviceNumberModelId: deviceNumberModelId,
		DeviceName:          reqDevice.Device.DeviceName,
		DeviceStatus:        reqDevice.Device.DeviceStatus,
		DeviceImages:        reqDevice.Device.DeviceImages,
		DeviceUseDesc:       reqDevice.Device.DeviceUseDesc,
		LaboratoryId:        reqDevice.Device.LaboratoryId,
		LaboratoryFloor:     reqDevice.Device.LaboratoryFloor,
		LaboratoryFloorId:   reqDevice.Device.LaboratoryFloorId,
	}
	err = dals_device.UpdateDevice(tx, modelDevice, 0)
	if err != nil {
		return resp, err
	}

	//3.记录日志
	beforeContent, _ := json.Marshal(modelDeviceList[0])
	nowContent, _ := json.Marshal(reqDevice.Device)
	modelLog := &models.MultiLog{
		Operator:            device2.LogUpdateDevice,
		MultiId:             reqDevice.MultiId,
		DeviceNumberModelId: deviceNumberModelId,
		BeforeContent:       string(beforeContent),
		AfterContent:        string(nowContent),
		OperateTime:         time.Unix(time.Now().Unix(), 0),
	}
	flag, err := dals_log.SetLog(tx, modelLog)
	if err != nil || !flag {
		return resp, err
	}

	resp.Message = device2.UpdateDeviceSuccessful
	return resp, nil
}

func (d *Device) MQueryDeviceByDNMIS(ctx context.Context, reqDevice *device.MQueryDeviceRequest) (*device.MQueryDeviceResponse, error) {
	var (
		resp           *device.MQueryDeviceResponse
		respDeviceList []*device.Device
	)
	resp = &device.MQueryDeviceResponse{}
	respDeviceList = make([]*device.Device, 0)
	resp.Message = device2.QueryDeviceFailed
	modelDeviceList, err := dals_device.QueryDeviceByDNMIS(nil, reqDevice.DeviceNumberModelId)
	if err != nil {
		return resp, err
	}
	if len(modelDeviceList) == 0 {
		resp.Message = device2.QueryDeviceNotExist
		return resp, device2.QueryDeviceNotExistErr
	}

	for _, modelDevice := range modelDeviceList {
		var respDevice = &device.Device{
			DeviceNumberId:      modelDevice.DeviceNumberId,
			DeviceModelId:       modelDevice.DeviceModelId,
			DeviceNumberModelId: modelDevice.DeviceNumberModelId,
			DeviceName:          modelDevice.DeviceName,
			DeviceStatus:        modelDevice.DeviceStatus,
			DeviceImages:        modelDevice.DeviceImages,
			DeviceUseDesc:       modelDevice.DeviceUseDesc,
			DeviceCreateTime:    modelDevice.DeviceCreateTime.Unix(),
			LaboratoryId:        modelDevice.LaboratoryId,
			LaboratoryFloor:     modelDevice.LaboratoryFloor,
			LaboratoryFloorId:   modelDevice.LaboratoryFloorId,
		}
		respDeviceList = append(respDeviceList, respDevice)
	}
	resp.Message = device2.QueryDeviceSuccessful
	resp.Device = respDeviceList
	return resp, nil
}
