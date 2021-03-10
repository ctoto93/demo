package mongo

import (
	"context"
	"log"
	"time"

	"github.com/ctoto93/demo"
	mapset "github.com/deckarep/golang-set"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	Id        primitive.ObjectID    `bson:"_id,omitempty"`
	CreatedAt time.Time             `bson:"created_at"`
	UpdatedAt time.Time             `bson:"updated_at"`
	Name      string                `bson:"name"`
	Credit    int                   `bson:"credit"`
	Students  *[]primitive.ObjectID `bson:"students"`
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

	studentIds := make([]primitive.ObjectID, 0)
	if len(dc.Students) > 0 {
		for _, ds := range dc.Students {
			oid, err := primitive.ObjectIDFromHex(ds.Id)
			if err != nil {
				return Course{}, err
			}

			studentIds = append(studentIds, oid)
		}
	}

	c.Students = &studentIds
	return c, nil
}

func (c *Course) toDemo() demo.Course {
	return demo.Course{
		Id:       c.Id.Hex(),
		Name:     c.Name,
		Credit:   c.Credit,
		Students: make([]demo.Student, 0),
	}
}

func (r *Repository) getCourses(oids []primitive.ObjectID) ([]demo.Course, error) {
	var demoCourses []demo.Course
	filter := bson.M{
		"_id": bson.M{
			"$in": oids,
		},
	}
	cur, err := r.Db.Collection("courses").Find(context.TODO(), filter)
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

func (r *Repository) getCourseByObjectId(oid primitive.ObjectID) (Course, error) {
	var c Course
	err := r.Db.Collection("courses").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&c)
	return c, err
}
func (r *Repository) GetCourse(id string) (demo.Course, error) {
	var c Course
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return demo.Course{}, err
	}
	err = r.Db.Collection("courses").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&c)
	if err != nil {
		return demo.Course{}, err
	}

	demoCourse := c.toDemo()

	if len(*c.Students) > 0 {
		demoStudents, err := r.getStudents(*c.Students)
		if err != nil {
			return demo.Course{}, err
		}
		demoCourse.Students = demoStudents
	}

	return demoCourse, nil
}

func (r *Repository) AddCourse(dc *demo.Course) error {
	c, err := newCourse(*dc)
	if err != nil {
		return err
	}

	res, err := r.Db.Collection("courses").InsertOne(context.TODO(), c)
	if err != nil {
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		c.Id = oid
	}

	dc.Id = c.Id.Hex()

	if err := r.addCourseForStudents(c.Id, *c.Students); err != nil {
		return err
	}

	return nil
}

func (r *Repository) EditCourse(dc *demo.Course) error {
	updatedCourse, err := newCourse(*dc)
	if err != nil {
		return err
	}

	currentCourse, err := r.getCourseByObjectId(updatedCourse.Id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", updatedCourse.Id}}
	update := bson.D{
		{"$set", updatedCourse},
	}
	_, err = r.Db.Collection("courses").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	addedStudents, removedStudents, err := differenceObjectIDs(*currentCourse.Students, *updatedCourse.Students)

	if len(addedStudents) > 0 {
		if err := r.addCourseForStudents(currentCourse.Id, addedStudents); err != nil {
			return err
		}
	}

	if len(removedStudents) > 0 {
		if err := r.removeCourseForStudents(currentCourse.Id, removedStudents); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) DeleteCourse(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.Db.Collection("courses").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (r *Repository) removeCourseForStudents(courseId primitive.ObjectID, oids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": oids,
		},
	}

	update := bson.M{
		"$pull": bson.M{
			"courses": courseId,
		},
	}
	_, err := r.Db.Collection("students").UpdateMany(context.Background(), filter, update)
	return err
}

func (r *Repository) addCourseForStudents(courseId primitive.ObjectID, oids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": oids,
		},
	}

	update := bson.M{
		"$push": bson.M{
			"courses": courseId,
		},
	}

	_, err := r.Db.Collection("students").UpdateMany(context.Background(), filter, update)
	return err
}

func toSliceObjectID(s []interface{}) ([]primitive.ObjectID, error) {
	var res []primitive.ObjectID
	for _, in := range s {
		if val, ok := in.(string); ok {
			oid, err := primitive.ObjectIDFromHex(val)
			if err != nil {
				return res, err
			}
			res = append(res, oid)
		} else {
			log.Fatal("Cannot convert")
		}

	}
	return res, nil
}

func toSliceInterface(s []primitive.ObjectID) []interface{} {
	var res []interface{}
	for _, in := range s {
		res = append(res, in.Hex())
	}
	return res
}

func toSliceString(s []primitive.ObjectID) []string {
	var res []string
	for _, oid := range s {
		res = append(res, oid.Hex())
	}
	return res
}

func differenceObjectIDs(current, updated []primitive.ObjectID) (added, removed []primitive.ObjectID, err error) {
	currentSet := mapset.NewSetFromSlice(toSliceInterface(current))
	updatedSet := mapset.NewSetFromSlice(toSliceInterface(updated))

	added, err = toSliceObjectID(updatedSet.Difference(currentSet).ToSlice())
	if err != nil {
		return
	}

	removed, err = toSliceObjectID(currentSet.Difference(updatedSet).ToSlice())
	if err != nil {
		return
	}

	return
}
