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
	Id        primitive.ObjectID    `bson:"_id,omitempty"`
	CreatedAt time.Time             `bson:"created_at"`
	UpdatedAt time.Time             `bson:"updated_at"`
	Name      string                `bson:"name"`
	Age       int                   `bson:"age"`
	Courses   *[]primitive.ObjectID `bson:"courses"`
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

	coursesIds := make([]primitive.ObjectID, 0)
	if len(ds.Courses) > 0 {
		for _, dc := range ds.Courses {
			oid, err := primitive.ObjectIDFromHex(dc.Id)
			if err != nil {
				return Student{}, err
			}

			coursesIds = append(coursesIds, oid)
		}
	}

	s.Courses = &coursesIds

	return s, nil
}

func (s *Student) toDemo() demo.Student {
	return demo.Student{
		Id:      s.Id.Hex(),
		Name:    s.Name,
		Age:     s.Age,
		Courses: make([]demo.Course, 0),
	}
}

func (r *repository) getStudents(oids []primitive.ObjectID) ([]demo.Student, error) {
	var demoStudents []demo.Student
	filter := bson.M{
		"_id": bson.M{
			"$in": oids,
		},
	}
	cur, err := r.db.Collection("students").Find(context.TODO(), filter)
	if err != nil {
		return []demo.Student{}, err
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s Student
		err := cur.Decode(&s)
		if err != nil {
			return []demo.Student{}, err
		}

		demoStudents = append(demoStudents, s.toDemo())
	}

	return demoStudents, nil
}

func (r *repository) GetStudent(id string) (demo.Student, error) {
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
	if len(*student.Courses) > 0 {
		demoStudent.Courses, err = r.getCourses(*student.Courses)
		if err != nil {
			return demo.Student{}, err
		}
	}

	return demoStudent, nil
}

func (r *repository) getStudentByObjectId(oid primitive.ObjectID) (Student, error) {
	var s Student
	err := r.db.Collection("students").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&s)
	return s, err
}

func (r *repository) removeStudentForCourses(studentId primitive.ObjectID, courseIds []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": courseIds,
		},
	}

	update := bson.M{
		"$pull": bson.M{
			"students": studentId,
		},
	}
	_, err := r.db.Collection("courses").UpdateMany(context.Background(), filter, update)
	return err
}

func (r *repository) addStudentForCourses(studentId primitive.ObjectID, courseIds []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": courseIds,
		},
	}

	update := bson.M{
		"$push": bson.M{
			"students": studentId,
		},
	}

	_, err := r.db.Collection("courses").UpdateMany(context.Background(), filter, update)
	return err
}

func (r *repository) AddStudent(ds *demo.Student) error {

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
	if err := r.addStudentForCourses(s.Id, *s.Courses); err != nil {
		return err
	}
	return nil
}

func (r *repository) EditStudent(ds *demo.Student) error {
	updatedStudent, err := newStudent(*ds)
	if err != nil {
		return err
	}

	currentStudent, err := r.getStudentByObjectId(updatedStudent.Id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", updatedStudent.Id}}
	update := bson.D{
		{"$set", updatedStudent},
	}
	_, err = r.db.Collection("students").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	addedStudents, removedStudents, err := differenceObjectIDs(*currentStudent.Courses, *updatedStudent.Courses)

	if len(addedStudents) > 0 {
		if err := r.addStudentForCourses(currentStudent.Id, addedStudents); err != nil {
			return err
		}
	}

	if len(removedStudents) > 0 {
		if err := r.removeStudentForCourses(currentStudent.Id, removedStudents); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) DeleteStudent(id string) error {
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
