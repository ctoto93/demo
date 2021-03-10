package factory

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/ctoto93/demo"
)

func NewStudent() demo.Student {
	return demo.Student{
		Name:    randomdata.FullName(randomdata.RandomGender),
		Age:     randomdata.Number(18, 25),
		Courses: make([]demo.Course, 0),
	}
}

func NewStudents(n int) []demo.Student {
	students := []demo.Student{}
	for i := 0; i < n; i++ {
		students = append(students, NewStudent())
	}

	return students
}
