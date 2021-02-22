package demo

import "time"

type Student struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Courses   []Course  `json:"courses,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Course struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Credit    int       `json:"credit"`
	Students  []Student `json:"students,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
