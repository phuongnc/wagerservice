package db

import (
	"log"
	"wagerservice/cmd/entity-server/db/model"
	"wagerservice/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) (*gorm.DB, func(), error) {
	db, err := gorm.Open(mysql.Open(cfg.DBConnection), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10000)

	cleanFunc := func() {
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Gorm db close error: %s", err)
		}
	}
	_ = MigrateTable(db)
	return db, cleanFunc, err
}

func MigrateTable(DB *gorm.DB) error {
	return DB.AutoMigrate(
		new(model.Wager),
		new(model.WagerTransaction),
	)
}
