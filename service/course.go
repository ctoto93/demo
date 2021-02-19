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
	repository courseRepository
}

func (cs *Course) Add(c *demo.Course) error {
	if len(c.Students) < MinNumOfStudents {
		return InsufficientStudentsErr
	}

	if len(c.Students) > MaxNumOfStudents {
		return ExceedingStudentsErr
	}
	return cs.repository.AddCourse(c)
}

func (cs *Course) Get(courseId int) (demo.Course, error) {
	return cs.repository.GetCourse(courseId)
}

func (cs *Course) Edit(c *demo.Course) error {
	return cs.repository.EditCourse(c)
}

func (cs *Course) Delete(courseId int) error {
	return cs.repository.DeleteCourse(courseId)
}
