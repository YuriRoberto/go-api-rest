package migrations

import (
	"github.com/YuriRoberto/go-api-rest/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
