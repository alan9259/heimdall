package platform

import (
	"fmt"
	"log"
	"os"

	"miu-auth-api-v1/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func New() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost password=WorkHappily123 user=postgres dbname=dbmessyitup sslmode=disable connect_timeout=30")
	//defer db.Close()
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open(
		"postgres",
		"host=localhost password=WorkHappily123 user=postgres dbname=dbmessyitup sslmode=disable connect_timeout=30")

	if err != nil {
		log.Fatalln("Can't connect to the db: ", err)
		fmt.Println("storage err: ", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Can't close the db connection: ", err)
		}
	}()

	//db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}

//TODO: err check
func AutoMigrate(db *gorm.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "miu." + defaultTableName
	}

	db.AutoMigrate(
		&model.Account{},
		&model.Location{},
	)
}
