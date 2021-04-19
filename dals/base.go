package dals

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB   = "mysql"
	Args = "root:12345678@/graduation_project?charset=utf8&parseTime=True&loc=Local"
)

func GetConn() (*gorm.DB, error) {
	db, err := gorm.Open(DB, Args)
	if err != nil {
		log.Printf("conn db is failed,err:%v\n", err)
		return nil, err
	}
	return db, nil
}
