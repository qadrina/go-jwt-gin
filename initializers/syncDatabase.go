package initializers

import "github.com/qadrina/go-jwt-gin/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
