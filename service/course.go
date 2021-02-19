package service

import "github.com/ctoto93/demo"

type Course struct {
	repository courseRepository
}

func (cs *Course) Add(c *demo.Course) error {
	return cs.repository.AddCourse(c)
}

func (cs *Course) Get(courseId int) (demo.Course, error) {
	return cs.repository.GetCourse(courseId)
}

func (cs *Course) Edit(c *demo.Course) error {
	return cs.repository.EditCourse(c)
}

func (cs *Course) Delete(courseId int) error {
	return cs.repository.DeleteCourse(courseId)
}
