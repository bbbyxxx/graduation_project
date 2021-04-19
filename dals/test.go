package dals

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`
	MemberNumber *string `gorm:"not null"`
	Num          int     `gorm:"AUTO_INCREMENT""`
	Address      string  `gorm:"index:addr"`
	IgnoreMe     int     `gorm:"-"`
}

type Animal struct {
	ID      int64
	Name    string `gorm:"default:'wocao'"`
	Age     string
	Xinzeng string
	wuqing  string
}

func main() {
	db, err := gorm.Open("mysql", "root:12345678@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db conn is failed.err:%v", err)
	}

	var users []User
	var user User
	//db.Table("user").Debug().Where("name = ?", "jinzhu").First(&user)
	//db.Table("user").Debug().Where("name in (?)", []string{"jinzhu", "wuqing"}).Find(&user)
	//db.Table("user").Where("name like ?", "%jin%").Debug().Find(&user)
	//db.Table("user").Debug().Where("name = ? and age >= ?", "jinzhu", 19).Find(&user)
	//db.Table("user").Debug().Where(&User{Name: "Jinzhu", Age: sql.NullInt64{
	//	Int64: 23,
	//	Valid: true,
	//}}).First(&user)
	//db.Table("user").Debug().Where(map[string]interface{}{"name": "jinzhu", "age": 20}).First(&user)
	//db.Table("user").Debug().Set("gorm:query_option", "for update").First(&user)
	//db.Table("user").Debug().Where("age > ?", db.Table("user").Select("age").Where("age = 23").SubQuery()).Find(&users)
	//rows, err := db.Table("user").Select("date(created_at) as date,sum(age) as total").Group("date(created_at)").Rows()
	//for rows.Next() {
	//	fmt.Println(rows)
	//}
	//user.Name = "无情啊"
	//user.Age = sql.NullInt64{
	//	Int64: 100,
	//	Valid: true,
	//}
	//user.Birthday = time.Now()
	//var s = "xxxcc"
	//user.MemberNumber = &s
	//user.Email = s

	db.Debug().Delete(&user).Where("id = ?", 22)

	fmt.Println(user)
	fmt.Println(users)

	//db.Table("user").AutoMigrate(&User{})

	//db.Table("animal").CreateTable(&Animal{})
	//var animal = Animal{Age: 99, Name: ""}
	//db.Table("animal").Create(&animal)

	//db.Table("user").CreateTable(User{})
	//for i := 0; i < 10; i++ {
	//	var tmp = fmt.Sprintf("%d%s", i, "xxx")
	//	it := &tmp
	//	user := User{Name: "Jinzhu", Age: sql.NullInt64{
	//		Int64: int64(18 + i),
	//		Valid: true,
	//	}, Email: *it, Birthday: time.Now(), MemberNumber: it}
	//	//
	//	//flag := db.NewRecord(user)
	//	//fmt.Println(flag)
	//	//
	//	db.Table("user").Create(&user)
	//}
	//
	//flag = db.NewRecord(user)
	//fmt.Println(flag)

	defer db.Close()
}
