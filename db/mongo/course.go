package mongo

import (
	"context"
	"time"

	"github.com/ctoto93/demo"
	"go.mongodb.org/mongo-driver/bson"
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
	c := Course{
		Name:   dc.Name,
		Credit: dc.Credit,
	}

	if dc.Id != "" {
		oid, err := primitive.ObjectIDFromHex(dc.Id)
		if err != nil {
			return Course{}, err
		}
		c.Id = oid
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
		students = append(students, s.toDemo())
	}

	return demo.Course{
		Id:       c.Id.Hex(),
		Name:     c.Name,
		Credit:   c.Credit,
		Students: students,
	}

}

func (c *Course) toDemo() demo.Course {
	return demo.Course{
		Id:     c.Id.Hex(),
		Name:   c.Name,
		Credit: c.Credit,
	}
}

func (r *Repository) getCourses(oids []primitive.ObjectID) ([]demo.Course, error) {
	var demoCourses []demo.Course
	filter := bson.M{
		"_id": bson.M{
			"$in": oids,
		},
	}
	cur, err := r.db.Collection("courses").Find(context.TODO(), filter)
	if err != nil {
		return []demo.Course{}, err
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var c Course
		err := cur.Decode(&c)
		if err != nil {
			return []demo.Course{}, err
		}

		demoCourses = append(demoCourses, c.toDemo())
	}

	return demoCourses, nil

}
