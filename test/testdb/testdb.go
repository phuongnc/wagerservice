package testdb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConnection = "admin:WagerService123@tcp(127.0.0.1:23306)/wager_test?charset=utf8mb4&parseTime=True&loc=Local"
)

func GetDB() (db *gorm.DB, release func()) {
	db, err := gorm.Open(mysql.Open(DBConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10000)

	cleanFunc := func() {
		err = sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}
	return db, cleanFunc
}
