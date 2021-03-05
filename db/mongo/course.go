package mongo

import (
	"context"
	"time"

	"github.com/ctoto93/demo"
	"github.com/davecgh/go-spew/spew"
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
	if len(c.Students) > 0 {
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

func (r *Repository) getCourseByObjectId(oid primitive.ObjectID) (Course, error) {
	return Course{}, nil
}

func (r *Repository) GetCourse(id string) (demo.Course, error) {
	var c Course
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return demo.Course{}, err
	}
	err = r.db.Collection("courses").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&c)
	if err != nil {
		return demo.Course{}, err
	}

	dc := c.toDemo()
	spew.Dump("xxxxx", c.Students)
	if len(c.Students) > 0 {
		demoStudents, err := r.getStudents(c.Students)
		if err != nil {
			return demo.Course{}, err
		}

		dc.Students = demoStudents
	}
	return dc, nil
}

func (r *Repository) AddCourse(c *demo.Course) error {
	return nil
}

func (r *Repository) EditCourse(c *demo.Course) error {
	return nil
}

func (r *Repository) DeleteCourse(id string) error {
	return nil
}
