package sql

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []Course `json:"courses,omitempty" gorm:"many2many:user_languages;"`
}
