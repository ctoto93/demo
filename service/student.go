package service

import (
	"errors"

	"github.com/ctoto93/demo"
)

const StudentMaxCredits = 30

var OverCreditErr = errors.New("Credit exceeding limit")

type Student struct {
	repository StudentRepository
}

func NewStudent(repo Repository) *Student {
	return &Student{repository: repo}
}

func isOverCredit(courses []demo.Course) bool {
	sum := 0
	for _, c := range courses {
		sum += c.Credit
	}

	return sum > StudentMaxCredits
}

func (ss *Student) Get(studentId string) (demo.Student, error) {
	return ss.repository.GetStudent(studentId)
}

func (ss *Student) Add(s *demo.Student) error {
	if isOverCredit(s.Courses) {
		return OverCreditErr
	}
	return ss.repository.AddStudent(s)
}

func (ss *Student) Edit(s *demo.Student) error {
	return ss.repository.EditStudent(s)
}

func (ss *Student) Delete(studentId string) error {
	return ss.repository.DeleteStudent(studentId)
}
