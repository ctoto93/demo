package mongo

import (
	"context"
	"log"
	"time"

	"github.com/ctoto93/demo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"name"`
	Age       int                `bson:"age"`
	Courses   []Course           `bson:"courses"`
}

func newStudentWithCourses(ds demo.Student) (Student, error) {
	s, err := newStudent(ds)
	if err != nil {
		return s, nil
	}

	var courses []Course
	if len(ds.Courses) > 0 {
		for _, dc := range ds.Courses {
			oid, err := primitive.ObjectIDFromHex(dc.Id)
			if err != nil {
				return Student{}, err
			}

			courses = append(courses, Course{
				Id:        oid,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Name:      dc.Name,
			})
		}
	}

	s.Courses = courses

	return s, nil
}

func newStudent(ds demo.Student) (Student, error) {
	s := Student{
		Name:      ds.Name,
		Age:       ds.Age,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if ds.Id != "" {
		oid, err := primitive.ObjectIDFromHex(ds.Id)
		if err != nil {
			return Student{}, err
		}
		s.Id = oid
	}

	return s, nil
}

func (s *Student) toDemoWithCourses() demo.Student {
	ds := s.toDemoWithoutCourses()

	var courses []demo.Course
	for _, c := range s.Courses {
		courses = append(courses, c.toDemoWithoutStudents())
	}

	ds.Courses = courses

	return ds
}

func (s *Student) toDemoWithoutCourses() demo.Student {
	return demo.Student{
		Id:   s.Id.Hex(),
		Name: s.Name,
		Age:  s.Age,
	}
}

func (s *Student) buildCoursesBsonArray() bson.A {
	var res bson.A
	for _, c := range s.Courses {
		res = append(res, c.Id)
	}

	return res
}

func (s *Student) buildStudentBsonDoc() bson.D {
	return bson.D{
		{"name", s.Name},
		{"age", s.Age},
		{"courses", s.buildCoursesBsonArray()},
	}
}

func (r *Repository) GetStudent(id string) (demo.Student, error) {
	var s Student
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return demo.Student{}, err
	}

	err = r.db.Collection("students").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&s)
	if err != nil {
		return demo.Student{}, err
	}

	for i, c := range s.Courses {
		err = r.db.Collection("courses").FindOne(context.TODO(), bson.M{"_id": c.Id}).Decode(&c)
		if err != nil {
			return demo.Student{}, err
		}
		s.Courses[i] = c
	}

	return s.toDemoWithCourses(), nil
}

func (r *Repository) AddStudent(ds *demo.Student) error {

	s, err := newStudentWithCourses(*ds)
	if err != nil {
		return err
	}

	res, err := r.db.Collection("students").InsertOne(context.TODO(), s.buildStudentBsonDoc())
	if err != nil {
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		s.Id = oid
	}

	ds.Id = s.Id.Hex()

	return nil
}

func (r *Repository) EditStudent(ds *demo.Student) error {
	s, err := newStudentWithCourses(*ds)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", s.Id}}
	update := bson.D{
		{"$set", s.buildStudentBsonDoc()},
	}
	_, err = r.db.Collection("students").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r *Repository) DeleteStudent(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.db.Collection("students").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
