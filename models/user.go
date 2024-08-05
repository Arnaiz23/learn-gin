package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
}
