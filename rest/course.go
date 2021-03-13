package rest

import (
	"github.com/ctoto93/demo"
	"github.com/gin-gonic/gin"
)

func (s *server) getCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	course, err := s.service.GetCourse(id)
	if err != nil {
		s.sendError(ctx, err)
		return
	}
	s.send(ctx, course)
}

func (s *server) addCourse(ctx *gin.Context) {
	var course demo.Course
	if err := ctx.BindJSON(&course); err != nil {
		s.sendError(ctx, err)
		return
	}

	if err := s.service.AddCourse(&course); err != nil {
		s.sendError(ctx, err)
		return
	}
	s.send(ctx, course)
}

func (s *server) editCourse(ctx *gin.Context) {
	var course demo.Course
	if err := ctx.BindJSON(&course); err != nil {
		s.sendError(ctx, err)
		return
	}

	if err := s.service.EditCourse(&course); err != nil {
		s.sendError(ctx, err)
		return
	}
	s.send(ctx, course)
}

func (s *server) deleteCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.service.DeleteCourse(id)
	if err != nil {
		s.sendError(ctx, err)
		return
	}
	s.send(ctx, nil)
}
