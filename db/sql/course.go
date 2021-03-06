package sql

import (
	"strconv"
	"time"

	"github.com/ctoto93/demo"
)

type Course struct {
	Id        int       `gorm:"primarykey"`
	Name      string    `json:"name"`
	Credit    int       `json:"credit"`
	Students  []Student `json:"courses" gorm:"many2many:student_courses"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Course) toDemo() demo.Course {
	dc := demo.Course{
		Id:       strconv.Itoa(c.Id),
		Name:     c.Name,
		Credit:   c.Credit,
		Students: make([]demo.Student, 0),
	}
	if len(c.Students) > 0 {
		for _, ds := range c.Students {
			dc.Students = append(dc.Students, ds.toDemo())
		}
	}

	return dc
}

func NewCourse(dc demo.Course) (Course, error) {
	c := Course{
		Name:   dc.Name,
		Credit: dc.Credit,
	}

	if dc.Id != "" {
		id, err := strconv.Atoi(dc.Id)
		if err != nil {
			return Course{}, err
		}
		c.Id = id
	}

	students := make([]Student, 0)

	if len(dc.Students) > 0 {
		for _, ds := range dc.Students {
			s, err := NewStudent(ds)
			if err != nil {
				return Course{}, err
			}
			students = append(students, s)
		}
	}

	c.Students = students

	return c, nil
}

func (r *repository) GetCourse(sid string) (demo.Course, error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return demo.Course{}, err
	}

	var c Course

	if err := r.db.Preload("Students").First(&c, id).Error; err != nil {
		return demo.Course{}, err
	}

	return c.toDemo(), nil

}

func (r *repository) AddCourse(dc *demo.Course) error {

	c, err := NewCourse(*dc)
	if err != nil {
		return err
	}

	if err := r.db.Create(&c).Error; err != nil {
		return err
	}

	dc.Id = strconv.Itoa(c.Id)

	return nil

}

func (r *repository) EditCourse(dc *demo.Course) error {

	c, err := NewCourse(*dc)
	if err != nil {
		return err
	}

	if err := r.db.Save(&c).Error; err != nil {
		return err
	}

	if err := r.db.Model(&c).Association("Students").Replace(c.Students); err != nil {
		return err
	}

	return nil

}

func (r *repository) DeleteCourse(sid string) error {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return err
	}
	return r.db.Delete(&Course{}, id).Error
}
