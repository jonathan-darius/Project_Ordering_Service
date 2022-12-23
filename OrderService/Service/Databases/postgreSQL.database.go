package Databases

import (
	"OrderService/Service/Configs"
	"OrderService/Service/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	DBClient *gorm.DB = dbInit()
)

func dbInit() *gorm.DB {
	dbURL := Configs.GetEnv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln("Cant Connect To Databases: ", err)
	}

	db.AutoMigrate(&Models.UserCart{}, &Models.Purchase{}, &Models.PurchaseDetail{})
	log.Println("DB Connected")
	return db
}
