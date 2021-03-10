package sql

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name     string   `json:"name"`
	Credit   int      `json:"credit"`
	Students []Course `json:"students,omitempty" gorm:"many2many:student_courses;"`
}
