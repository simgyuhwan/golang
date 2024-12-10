package migration

import (
	"gofiber/database"
	"gofiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
}