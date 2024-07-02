package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"jhphon0730/DevConnector/models"
	"log"
)

// type DB struct *gorm.DB ( var DB *gorm.DB ) -> DB is a pointer to gorm.DB

var (
	DB *gorm.DB
	err error
)

func InitialDatabase() {
	dsn := "host=localhost user=jhkim password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Auto migrate
	DB.AutoMigrate(
		// User
		&models.User{}, &models.Experience{}, &models.Education{},
	)

	getTableList()
}

func getTableList() {
	tables := []string{}
	DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)

	log.Println("Tables in 'postgres' database:")
	for _, table := range tables {
		log.Println("TABLE: ", table)
	}
}
