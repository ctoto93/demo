package sql

import (
	"strconv"
	"time"

	"github.com/ctoto93/demo"
)

type Student struct {
	Id        int       `gorm:"primarykey"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Courses   *[]Course `json:"courses,omitempty" gorm:"many2many:user_languages;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newStudent(ds demo.Student) (Student, error) {
	s := Student{
		Name: ds.Name,
		Age:  ds.Age,
	}

	if ds.Id != "" {
		id, err := strconv.Atoi(ds.Id)
		if err != nil {
			return Student{}, err
		}
		s.Id = id
	}

	courses := make([]Course, 0)
	s.Courses = &courses

	return s, nil
}

func (s *Student) toDemo() demo.Student {
	return demo.Student{
		Id:      strconv.Itoa(s.Id),
		Name:    s.Name,
		Age:     s.Age,
		Courses: make([]demo.Course, 0),
	}
}

func (r *repository) GetStudent(sid string) (demo.Student, error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return demo.Student{}, err
	}

	var s Student

	if err := r.db.First(&s, id).Error; err != nil {
		return demo.Student{}, err
	}

	return s.toDemo(), nil

}

func (r *repository) AddStudent(ds *demo.Student) error {

	s, err := newStudent(*ds)
	if err != nil {
		return err
	}

	if err := r.db.Create(&s).Error; err != nil {
		return err
	}
	ds.Id = strconv.Itoa(s.Id)

	return nil

}

func (r *repository) DeleteStudent(sid string) error {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return err
	}
	return r.db.Delete(&Student{}, id).Error
}
