package dal_person

import (
	"lab_device_management_person/dals"
	"lab_device_management_person/models"
	"lab_device_management_person/service"

	"github.com/jinzhu/gorm"
)

const (
	PersonTableName = "person"
)

func LoginValid(db *gorm.DB, multiId string, password string) ([]*models.Person, error) {
	var (
		res []*models.Person
		err error
	)
	if db == nil {
		db, _ = dals.GetConn()
	}

	err = db.Table(PersonTableName).Where("multi_id = ? and password = ? and is_deleted = 0", multiId, password).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryPersonByMultiIdAndIndentity(db *gorm.DB, person *models.Person) ([]*models.Person, error) {
	var (
		res []*models.Person
		err error
	)
	if db == nil {
		db, _ = dals.GetConn()
	}
	err = db.Table(PersonTableName).Where("multi_id = ? and indentity = ?", person.MultiId, person.Indentity).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddPerson(db *gorm.DB, person *models.Person) (bool, error) {
	var (
		err error
	)
	if db == nil {
		db, _ = dals.GetConn()
	}
	err = db.Table(PersonTableName).Create(person).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdatePerson(db *gorm.DB, person *models.Person, isDelete bool) (bool, error) {
	var (
		err error
	)
	if db == nil {
		db, err = dals.GetConn()
	}
	//删除请求
	if isDelete {
		person.IsDeleted = 1
	}
	argsMap := service.UpdateModelPerson2Map(person)
	err = db.Table(PersonTableName).Where("multi_id = ?", person.MultiId).Update(argsMap).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateLoginTimePerson(db *gorm.DB, person *models.Person, isDelete bool) (bool, error) {
	var (
		err error
	)
	if db == nil {
		db, err = dals.GetConn()
	}

	err = db.Table(PersonTableName).Select("last_login_time, login_time").Where("multi_id = ?", person.MultiId).Updates(map[string]interface{}{"last_login_time": person.LastLoginTime, "login_time": person.LoginTime}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
