package service

import "github.com/ctoto93/demo"

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
