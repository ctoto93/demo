package service

import (
	"errors"

	"github.com/ctoto93/demo"
)

var UnimplementedMethodErr = errors.New("Not yet implemented")

type StudentRepository interface {
	GetStudent(id string) (demo.Student, error)
	AddStudent(s *demo.Student) error
	EditStudent(s *demo.Student) error
	DeleteStudent(id string) error
}
type CourseRepository interface {
	GetCourse(id string) (demo.Course, error)
	AddCourse(c *demo.Course) error
	EditCourse(c *demo.Course) error
	DeleteCourse(id string) error
}

type Repository interface {
	StudentRepository
	CourseRepository
}

type UnimplementedRepository struct{}

func (*UnimplementedRepository) GetStudent(id string) (demo.Student, error) {
	return demo.Student{}, UnimplementedMethodErr
}

func (*UnimplementedRepository) AddStudent(s *demo.Student) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) EditStudent(s *demo.Student) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) DeleteStudent(id string) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) GetCourse(id string) (demo.Course, error) {
	return demo.Course{}, UnimplementedMethodErr
}

func (*UnimplementedRepository) AddCourse(s *demo.Course) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) DeleteCourse(id string) error {
	return UnimplementedMethodErr
}
