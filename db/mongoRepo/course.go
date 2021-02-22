package mongoRepo

import (
	"time"

	"github.com/ctoto93/demo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"name"`
	Credit    int                `bson:"credit"`
	Students  []Student          `bson:"students"`
}

func newCourse(dc demo.Course) (Course, error) {
	oid, err := primitive.ObjectIDFromHex(dc.Id)
	if err != nil {
		return Course{}, err
	}

	c := Course{
		Id:     oid,
		Name:   dc.Name,
		Credit: dc.Credit,
	}

	return c, nil
}

func newCourseWithStudents(dc demo.Course) (Course, error) {
	c, err := newCourse(dc)
	if err != nil {
		return c, err
	}

	var students []Student
	for _, ds := range dc.Students {
		s, err := newStudent(ds)
		if err != nil {
			return Course{}, err
		}
		students = append(students, s)
	}
	c.Students = students
	return c, nil
}

func (c *Course) toDemoWithStudents() demo.Course {
	var students []demo.Student

	for _, s := range c.Students {
		students = append(students, s.toDemoWithoutCourses())
	}

	return demo.Course{
		Id:       c.Id.Hex(),
		Name:     c.Name,
		Credit:   c.Credit,
		Students: students,
	}

}

func (c *Course) toDemoWithoutStudents() demo.Course {
	return demo.Course{
		Id:     c.Id.Hex(),
		Name:   c.Name,
		Credit: c.Credit,
	}
}
