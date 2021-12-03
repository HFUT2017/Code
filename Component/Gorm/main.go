package main

import (
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id         int64
	UserName   string `gorm:"column:username"`
	PassWord   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

func main() {
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := "3306"
	dbname := "gorm"
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("连接数据库失败!,error=%v\n", err)
	}
	user := User{
		UserName:   "fanyinghao",
		PassWord:   "123456",
		CreateTime: time.Now().Unix(),
	}
	if err := db.Create(&user).Error; err != nil {
		log.Errorf("插入失败!err=%v\n", err)
		return
	}
	user = User{}
	result := db.Where("username = ?", "fanyinghao").First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("找不到记录\n")
		return
	}
	fmt.Println(user.UserName, user.PassWord)
}
