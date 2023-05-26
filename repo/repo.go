package repo

import (
	"JWT-authentication/database"
	"JWT-authentication/models"

	"fmt"
)

func CreatedUserRecord(user models.User) (models.User, error) {
	result := database.Db.Create(&user)
	if result.Error != nil {
		fmt.Printf("result.Error: %v\n", result.Error)
		return models.User{}, result.Error
	}
	return user, nil
}
