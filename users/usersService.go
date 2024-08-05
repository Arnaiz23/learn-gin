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

// TODO: Return only the fields that I want (delete `DeletedAt`)
func GetUserS(id int) (*models.User, error) {
  var user *models.User

  result := database.DB.First(&user, id)

  if result.Error != nil {
		return nil, errors.New("This user doesn't exists")
  }

  return user, nil
}
