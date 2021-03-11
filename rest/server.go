package rest

import (
	"github.com/ctoto93/demo"
	"github.com/gin-gonic/gin"
)

type server struct {
	router  *gin.Engine
	service *demo.Service
}

func (s *server) initRouter() {
	v1 := s.router.Group("/v1")
	{
		student := v1.Group("/students")
		{
			student.GET("/:id", s.getStudent)
			student.POST("/:id", s.addStudent)
			student.PUT("/:id", s.editStudent)
			student.DELETE("/:id", s.deleteStudent)
		}

		course := v1.Group("/courses")
		{
			course.GET("/:id", s.getCourse)
			course.POST("/:id", s.addCourse)
			course.PUT("/:id", s.editCourse)
			course.DELETE("/:id", s.deleteCourse)
		}
	}
}

func (s *server) Serve(address string) error {
	return s.router.Run(address)
}

func NewServer(repo demo.Repository) *server {
	s := &server{
		router:  gin.Default(),
		service: demo.NewService(repo),
	}

	s.initRouter()

	return s
}
