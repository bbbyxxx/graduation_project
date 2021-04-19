package service

import "lab_device_management_device/models"

func ModelDevice2Args(modelDevice *models.Device) map[string]interface{} {
	return map[string]interface{}{
		"device_number_id":       modelDevice.DeviceNumberId,
		"device_model_id":        modelDevice.DeviceModelId,
		"device_number_model_id": modelDevice.DeviceNumberModelId,
		"device_name":            modelDevice.DeviceName,
		"device_status":          modelDevice.DeviceStatus,
		"device_images":          modelDevice.DeviceImages,
		"device_use_desc":        modelDevice.DeviceUseDesc,
		"laboratory_id":          modelDevice.LaboratoryId,
		"laboratory_floor":       modelDevice.LaboratoryFloor,
		"laboratory_floor_id":    modelDevice.LaboratoryFloorId,
		"is_deleted":             modelDevice.IsDeleted,
	}
}
