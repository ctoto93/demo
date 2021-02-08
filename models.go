package demo

type Student struct {
	Id int
	Name    string
	Age     int
	Courses []Course
}

type Course struct {
	Id       int
	Name     string
	Credit   int
	Students []Student
}
