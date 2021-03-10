package sql

import (
	"strconv"
	"time"

	"github.com/ctoto93/demo"
	"github.com/mitchellh/mapstructure"
)

type Course struct {
	Id        int        `gorm:"primarykey"`
	Name      string     `json:"name"`
	Credit    int        `json:"credit"`
	Students  *[]Student `json:"courses,omitempty" gorm:"many2many:course_courses;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Course) toDemo() demo.Course {
	return demo.Course{
		Id:       strconv.Itoa(c.Id),
		Name:     c.Name,
		Credit:   c.Credit,
		Students: make([]demo.Student, 0),
	}
}

func newCourse(in interface{}) (Course, error) {
	var c Course
	if err := mapstructure.Decode(in, &c); err != nil {
		return Course{}, err
	}

	students := make([]Student, 0)
	c.Students = &students

	return c, nil
}

func (r *repository) GetCourse(sid string) (demo.Course, error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return demo.Course{}, err
	}

	var c Course

	if err := r.db.First(&c, id).Error; err != nil {
		return demo.Course{}, err
	}

	return c.toDemo(), nil

}

func (r *repository) AddCourse(dc *demo.Course) error {

	c, err := newCourse(dc)
	if err != nil {
		return err
	}

	if err := r.db.Create(&c).Error; err != nil {
		return err
	}

	dc.Id = strconv.Itoa(c.Id)

	return nil

}
