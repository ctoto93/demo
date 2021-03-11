package demo

import (
	"errors"
	"fmt"
)

var (
	OverCreditErr           = errors.New("Credit exceeding limit")
	UnimplementedMethodErr  = errors.New("Not yet implemented")
	InsufficientStudentsErr = fmt.Errorf("A course should have min %v students", MinNumOfStudents)
	ExceedingStudentsErr    = fmt.Errorf("A course should have max %v students", MaxNumOfStudents)
)

const (
	StudentMaxCredits = 30
	MinNumOfStudents  = 5
	MaxNumOfStudents  = 30
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (serv *Service) AddCourse(c *Course) error {
	if err := c.ValidateCourseRequirement(); err != nil {
		return err
	}
	return serv.repo.AddCourse(c)
}

func (cs *Service) GetCourse(courseId string) (Course, error) {
	return cs.repo.GetCourse(courseId)
}

func (cs *Service) EditCourse(c *Course) error {
	if err := c.ValidateCourseRequirement(); err != nil {
		return err
	}
	return cs.repo.EditCourse(c)
}

func (cs *Service) DeleteCourse(courseId string) error {
	return cs.repo.DeleteCourse(courseId)
}

func (serv *Service) GetStudent(studentId string) (Student, error) {
	return serv.repo.GetStudent(studentId)
}

func (serv *Service) AddStudent(s *Student) error {
	if s.IsOverCredit() {
		return OverCreditErr
	}
	return serv.repo.AddStudent(s)
}

func (serv *Service) EditStudent(s *Student) error {
	if s.IsOverCredit() {
		return OverCreditErr
	}
	return serv.repo.EditStudent(s)
}

func (serv *Service) DeleteStudent(studentId string) error {
	return serv.repo.DeleteStudent(studentId)
}
