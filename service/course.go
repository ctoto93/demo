package service

import (
	"fmt"

	"github.com/ctoto93/demo"
)

const (
	MinNumOfStudents = 5
	MaxNumOfStudents = 30
)

var (
	InsufficientStudentsErr = fmt.Errorf("A course should have min %v students", MinNumOfStudents)
	ExceedingStudentsErr    = fmt.Errorf("A course should have max %v students", MaxNumOfStudents)
)

type Course struct {
	repository CourseRepository
}

func NewCourse(repo Repository) *Course {
	return &Course{repository: repo}
}

func validateCourseRequirement(c demo.Course) error {
	numStudents := len(c.Students)
	if numStudents < MinNumOfStudents {
		return InsufficientStudentsErr
	}

	if numStudents > MaxNumOfStudents {
		return ExceedingStudentsErr
	}

	return nil
}

func (cs *Course) Add(c *demo.Course) error {
	if err := validateCourseRequirement(*c); err != nil {
		return err
	}
	return cs.repository.AddCourse(c)
}

func (cs *Course) Get(courseId string) (demo.Course, error) {
	return cs.repository.GetCourse(courseId)
}

func (cs *Course) Edit(c *demo.Course) error {
	if err := validateCourseRequirement(*c); err != nil {
		return err
	}
	return cs.repository.EditCourse(c)
}

func (cs *Course) Delete(courseId string) error {
	return cs.repository.DeleteCourse(courseId)
}
