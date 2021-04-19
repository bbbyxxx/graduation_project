package rpc

import (
	"context"
	"lab_device_management_api/model"
	"lab_device_management_api/proto/device/device"
	"log"

	"google.golang.org/grpc"
)

const (
	AddressDevice = "127.0.0.1:8002"
)

func UpdateDevice(ctx context.Context, conn *grpc.ClientConn, modelDevice *model.Device, multiId string) (*device.UpdateDeviceResponse, error) {
	var (
		err  error
		resp *device.UpdateDeviceResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(AddressDevice, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[UpdateDevice] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := device.NewDeviceClient(conn)
	if modelDevice.IsDeleted == 0 {
		deviceReq := &device.Device{
			DeviceNumberId:      modelDevice.DeviceNumberId,
			DeviceModelId:       modelDevice.DeviceNumberId,
			DeviceNumberModelId: modelDevice.DeviceNumberModelId,
			DeviceName:          modelDevice.DeviceName,
			DeviceStatus:        modelDevice.DeviceStatus,
			DeviceImages:        modelDevice.DeviceImages,
			DeviceUseDesc:       modelDevice.DeviceUseDesc,
			LaboratoryId:        modelDevice.LaboratoryId,
			LaboratoryFloor:     modelDevice.LaboratoryFloor,
			LaboratoryFloorId:   modelDevice.LaboratoryFloorId,
		}
		req := &device.UpdateDeviceRequest{
			Device:  deviceReq,
			MultiId: multiId,
		}
		resp, err = c.UpdateDevice(ctx, req)
		if err != nil {
			log.Printf("[UpdateDevice] call rpc UpdateDevice is failed,err:%v\n", err)
			return resp, err
		}
	} else {
		req := &device.DeleteDeviceRequest{
			DeviceNumberModelId: modelDevice.DeviceNumberModelId,
			MultiId:             multiId,
			IsDeleted:           1,
		}
		respDelete, err := c.DeleteDevice(ctx, req)
		resp = &device.UpdateDeviceResponse{
			Message: respDelete.Message,
		}
		if err != nil {
			log.Printf("[UpdateDevice] call rpc DeleteDevice is failed,err:%v\n", err)
			return resp, err
		}
	}
	return resp, nil
}

func AddDevice(ctx context.Context, conn *grpc.ClientConn, modelDevice *model.Device, multiId string) (*device.AddDeviceResponse, error) {
	var (
		err  error
		resp *device.AddDeviceResponse
	)
	if conn == nil {
		conn, err = grpc.Dial(AddressDevice, grpc.WithInsecure())
	}
	if err != nil {
		log.Printf("[Login] Dial is failed,err:%v\n", err)
		return resp, err
	}
	c := device.NewDeviceClient(conn)
	deviceReq := &device.Device{
		DeviceNumberId:      modelDevice.DeviceNumberId,
		DeviceModelId:       modelDevice.DeviceNumberId,
		DeviceNumberModelId: modelDevice.DeviceNumberModelId,
		DeviceName:          modelDevice.DeviceName,
		DeviceStatus:        modelDevice.DeviceStatus,
		DeviceImages:        modelDevice.DeviceImages,
		DeviceUseDesc:       modelDevice.DeviceUseDesc,
		LaboratoryId:        modelDevice.LaboratoryId,
		LaboratoryFloor:     modelDevice.LaboratoryFloor,
		LaboratoryFloorId:   modelDevice.LaboratoryFloorId,
	}
	req := &device.AddDeviceRequest{
		Device:  deviceReq,
		MultiId: multiId,
	}
	resp, err = c.AddDevice(ctx, req)
	if err != nil {
		log.Printf("[Login] call rpc Login is failed,err:%v\n", err)
		return resp, err
	}
	return resp, nil
}
