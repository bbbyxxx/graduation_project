package service

import (
	"lab_device_management_person/models"
)

func UpdateModelPerson2Map(p *models.Person) map[string]interface{} {
	field2Value := make(map[string]interface{})
	field2Value["multi_id"] = p.MultiId
	field2Value["name"] = p.Name
	field2Value["class"] = p.Class
	field2Value["grade"] = p.Grade
	field2Value["indentity"] = p.Indentity
	field2Value["major"] = p.Major
	field2Value["password"] = p.Password
	field2Value["phone"] = p.Phone
	field2Value["sex"] = p.Sex
	field2Value["update_time"] = p.UpdateTime
	field2Value["is_deleted"] = p.IsDeleted
	return field2Value
}
