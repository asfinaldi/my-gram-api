package database

import (
	"fmt"
	"log"
	"final-project/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	var (
		env      = os.Getenv("ENV")
		host     = os.Getenv("PGHOST")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDBNAME")
		port     = os.Getenv("PGPORT")
		dsn      = ""
		db       *gorm.DB
		err      error
	)

	if env == "production" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	}

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true}); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{}); err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

	return db
}
