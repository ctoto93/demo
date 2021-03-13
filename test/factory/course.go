package factory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/ctoto93/demo"
)

func NewCourse() demo.Course {
	c := demo.Course{
		Name:     randomdata.Adjective(),
		Credit:   1,
		Students: make([]demo.Student, 0),
	}
	return c
}

func NewCourseWithStudents(numStudents int) demo.Course {
	c := NewCourse()

	var students []demo.Student
	for i := 0; i < numStudents; i++ {
		s := NewStudent()
		students = append(students, s)
	}
	c.Students = students
	return c
}

func NewCourses(n int) []demo.Course {
	courses := []demo.Course{}
	for i := 0; i < n; i++ {
		courses = append(courses, NewCourse())
	}

	return courses
}

func NewCourseWithId() demo.Course {
	s := NewCourse()
	s.Id = randomdata.Digits(3)
	return s
}
