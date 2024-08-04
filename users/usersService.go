package users

import (
	"errors"

	"ginserver/database"
)

func GetUsersS() ([]User, error) {

	rows, err := database.DB.Query("SELECT id, name, last_name, age FROM users")

	if err != nil {
		return nil, errors.New("Users not found")
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Age)
		if err != nil {
			return nil, errors.New("errorr???")
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserS(id int) (*User, error) {
	var foundUser *User
	var Users []User

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
