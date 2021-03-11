package demo

type Repository interface {
	GetStudent(id string) (Student, error)
	AddStudent(s *Student) error
	EditStudent(s *Student) error
	DeleteStudent(id string) error
	GetCourse(id string) (Course, error)
	AddCourse(c *Course) error
	EditCourse(c *Course) error
	DeleteCourse(id string) error
}

type UnimplementedRepository struct{}

func (*UnimplementedRepository) GetStudent(id string) (Student, error) {
	return Student{}, UnimplementedMethodErr
}

func (*UnimplementedRepository) AddStudent(s Student) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) EditStudent(s Student) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) DeleteStudent(id string) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) GetCourse(id string) (Course, error) {
	return Course{}, UnimplementedMethodErr
}

func (*UnimplementedRepository) AddCourse(s Course) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) EditCourse(s Course) error {
	return UnimplementedMethodErr
}

func (*UnimplementedRepository) DeleteCourse(id string) error {
	return UnimplementedMethodErr
}
