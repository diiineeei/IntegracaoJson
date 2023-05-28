package models

import (
	"gorm.io/gorm"
)

type Contacts struct {
	gorm.Model
	Name      string `json:"name"`
	CellPhone string `json:"cellphone"`
}
