package dals

import (
	"log"

	"github.com/jinzhu/gorm"
)

const (
	DB   = "mysql"
	Args = "root:12345678@/graduation_project?charset=utf8&parseTime=True&loc=Local"
)

func GetConn() (*gorm.DB, error) {
	db, err := gorm.Open(DB, Args)
	if err != nil {
		log.Println("conn db is failed,err:%v", err)
		return nil, err
	}
	return db, nil
}
