package demo

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type Student struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Courses   []Course  `json:"courses,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Course struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Credit    int       `json:"credit"`
	Students  []Student `json:"students,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func ToStudent(in interface{}) (Student, error) {
	var s Student
	err := mapstructure.Decode(in, &s)
	return s, err
}

func ToCourse(in interface{}) (Course, error) {
	var c Course
	err := mapstructure.Decode(in, &c)
	return c, err
}
