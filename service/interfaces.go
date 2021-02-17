package service

import "github.com/ctoto93/demo"

type studentRepository interface {
	GetStudent(id int) (demo.Student, error)
	AddStudent(s *demo.Student) error
	EditStudent(s *demo.Student) error
	DeleteStudent(id int) error
}
type courseRepository interface {
	GetCourse(id int) (demo.Course, error)
	AddCourse(c *demo.Course) error
	EditCourse(c *demo.Course) error
	DeleteCourse(id int) error
}
