package rest

import (
	"net/http"

	"github.com/ctoto93/demo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
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
			student.POST("/", s.addStudent)
			student.PUT("/:id", s.editStudent)
			student.DELETE("/:id", s.deleteStudent)
		}

		course := v1.Group("/courses")
		{
			course.GET("/:id", s.getCourse)
			course.POST("/", s.addCourse)
			course.PUT("/:id", s.editCourse)
			course.DELETE("/:id", s.deleteCourse)
		}
	}
}

func (s *server) sendError(ctx *gin.Context, err error) {
	var code int
	switch err {
	case gorm.ErrRecordNotFound, mongo.ErrNoDocuments:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	r := Response{
		Meta: MetaResp{
			Success: false,
			Error:   err.Error(),
		},
	}
	ctx.JSON(code, r)
}

func (s *server) send(ctx *gin.Context, data interface{}) {
	r := Response{
		Meta: MetaResp{
			Success: true,
		},
		Data: data,
	}
	ctx.JSON(http.StatusOK, r)
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

type MetaResp struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type Response struct {
	Meta MetaResp    `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
