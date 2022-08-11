package commons

import (
	"log"

	"github.com/Jonagarme/CRUD-SQL/models"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:0801@/test")

	if error != nil {
		log.Fatal(error)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando...")
	db.AutoMigrate(&models.Persona{})
}
