package platform

import (
	"fmt"
	"log"
	"os"

	"heimdall/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New() *gorm.DB {
	dbDSN := os.Getenv("MIU_DSN")
	if dbDSN == "" {
		dbDSN = "host=localhost user=postgres password=WorkHappily123 dbname=dbmessyitup sslmode=disable connect_timeout=30"
	}
	db, err := gorm.Open(
		"postgres",
		dbDSN)

	if err != nil {
		log.Fatalln("Can't connect to the db: ", err)
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func InitTestDB() *gorm.DB {
	dbDSN := os.Getenv("MIU_DSN")
	if dbDSN == "" {
		dbDSN = "host=localhost user=postgres password=WorkHappily123 db.name=dbmessyitup_test sslmode=disable connect_timeout=30"
	}
	db, err := gorm.Open(
		"postgres",
		dbDSN)

	if err != nil {
		log.Fatalln("Can't connect to the db: ", err)
		fmt.Println("storage err: ", err)
	}

	// defer func() {
	// 	if err := db.Close(); err != nil {
	// 		log.Fatalln("Can't close the db connection: ", err)
	// 	}
	// }()

	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	// if err := os.Remove("./../realworld_test.db"); err != nil {
	// 	return err
	// }
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "miu." + defaultTableName
	}

	db.AutoMigrate(
		&model.Gender{},
		&model.Device{},
		&model.Location{},
		&model.Account{},
		&model.Config{},
		&model.RevokedToken{},
		&model.Pin{},
	)
}
