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

func (s *Student) HasCourse(c Course) bool {
	for i := range s.Courses {
		if s.Courses[i].Id == c.Id {
			return true
		}
	}

	return false
}

func (s *Student) IsOverCredit() bool {
	sum := 0
	for _, c := range s.Courses {
		sum += c.Credit
	}

	return sum > StudentMaxCredits
}

type Course struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Credit    int       `json:"credit"`
	Students  []Student `json:"students,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (c *Course) HasStudent(s Student) bool {
	for i := range c.Students {
		if c.Students[i].Id == s.Id {
			return true
		}
	}

	return false
}

func (c *Course) ValidateCourseRequirement() error {
	numStudents := len(c.Students)
	if numStudents < MinNumOfStudents {
		return InsufficientStudentsErr
	}

	if numStudents > MaxNumOfStudents {
		return ExceedingStudentsErr
	}

	return nil
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