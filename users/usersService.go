package users

import "errors"

func GetUsersS() ([]User, error) {
	if Users == nil {
		return nil, errors.New("Users not found")
	}

  return Users, nil
}

func GetUserS (id int) (*User, error) {
	var foundUser *User

	for _, user := range Users {
		if user.Id == id {
			foundUser = &user
		}
	}

	if foundUser == nil {
    return nil, errors.New("This user doesn't exists")
	}

  return foundUser, nil
}
