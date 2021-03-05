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
	Id        primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"`
	Name      string               `bson:"name"`
	Age       int                  `bson:"age"`
	Courses   []primitive.ObjectID `bson:"courses"`
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

	var courseIds []primitive.ObjectID
	if len(ds.Courses) > 0 {
		for _, dc := range ds.Courses {
			oid, err := primitive.ObjectIDFromHex(dc.Id)
			if err != nil {
				return Student{}, err
			}

			courseIds = append(courseIds, oid)
		}
	}

	s.Courses = courseIds

	return s, nil
}

func (s *Student) toDemo() demo.Student {
	return demo.Student{
		Id:   s.Id.Hex(),
		Name: s.Name,
		Age:  s.Age,
	}
}

func (r *Repository) GetStudent(id string) (demo.Student, error) {
	var student Student
	var demoStudent demo.Student

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return demo.Student{}, err
	}

	student, err = r.getStudentByObjectId(oid)
	if err != nil {
		return demo.Student{}, err
	}

	demoStudent = student.toDemo()
	if len(student.Courses) > 0 {
		demoStudent.Courses, err = r.getCourses(student.Courses)
		if err != nil {
			return demo.Student{}, err
		}
	}

	return demoStudent, nil
}

func (r *Repository) getStudentByObjectId(oid primitive.ObjectID) (Student, error) {
	var s Student
	err := r.db.Collection("students").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&s)
	return s, err
}

func (r *Repository) AddStudent(ds *demo.Student) error {

	s, err := newStudent(*ds)
	if err != nil {
		return err
	}

	res, err := r.db.Collection("students").InsertOne(context.TODO(), s)
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
	s, err := newStudent(*ds)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", s.Id}}
	update := bson.D{
		{"$set", s},
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
