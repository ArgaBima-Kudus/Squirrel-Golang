package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsnmysqldeploy := "u472991515_vincenntes:vincen241!Aa@tcp(https://auth-db1031.hstgr.io/)/u472991515_dbceria?charset=utf8mb4&parseTime=True&loc=Local"
	// dsnpsql := "host=localhost user=postgres password=123456 dbname=dbceria port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// MYSQL
	// dsnmysql := "root:123456@tcp(localhost:3307)/dbceria?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsnmysql), &gorm.Config{NamingStrategy: schema.NamingStrategy{
	// SingularTable: true,
	// }})

	// POSTGRES
	dsn := "host=localhost user=postgres password=12345 dbname=dbceria port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {
		fmt.Printf("Error")
		return
	}
	DB = db
}
