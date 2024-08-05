package users

import (
	"errors"

	"ginserver/database"
	"ginserver/models"
)

func GetUsersS() ([]models.User, error) {

  var users []models.User
  result := database.DB.Find(&users)


	if result.Error != nil {
		return nil, errors.New("Users not found")
	}

	return users, nil
}

// TODO
func GetUserS(id int) (*models.User, error) {
	var foundUser *models.User
	var Users []models.User

	for _, user := range Users {
		if user.ID == id {
			foundUser = &user
		}
	}

	if foundUser == nil {
		return nil, errors.New("This user doesn't exists")
	}

	return foundUser, nil
}
