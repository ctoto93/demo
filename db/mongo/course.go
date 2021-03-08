package mongo

import (
	"context"
	"log"
	"time"

	"github.com/ctoto93/demo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	Id        primitive.ObjectID   `bson:"_id,omitempty"`
	CreatedAt time.Time            `bson:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at"`
	Name      string               `bson:"name"`
	Credit    int                  `bson:"credit"`
	Students  []primitive.ObjectID `bson:"students"`
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

	var studentsIds []primitive.ObjectID
	if len(dc.Students) > 0 {
		for _, ds := range dc.Students {
			oid, err := primitive.ObjectIDFromHex(ds.Id)
			if err != nil {
				return Course{}, err
			}

			studentsIds = append(studentsIds, oid)
		}
	}

	c.Students = studentsIds
	return c, nil
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

	if len(c.Students) > 0 {
		demoStudents, err := r.getStudents(c.Students)
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

	return nil
}

func (r *Repository) EditCourse(dc *demo.Course) error {
	c, err := newCourse(*dc)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", c.Id}}
	update := bson.D{
		{"$set", c},
	}
	_, err = r.Db.Collection("courses").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
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
