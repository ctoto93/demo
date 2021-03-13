package rest

import (
	"github.com/ctoto93/demo"
	"github.com/gin-gonic/gin"
)

func (s *server) getStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	student, err := s.service.GetStudent(id)
	if err != nil {
		s.sendError(ctx, err)
		return
	}
	s.send(ctx, student)
}

func (s *server) addStudent(ctx *gin.Context) {
	var student demo.Student
	if err := ctx.BindJSON(&student); err != nil {
		s.sendError(ctx, err)
		return
	}

	if err := s.service.AddStudent(&student); err != nil {
		return
	}
	s.send(ctx, student)
}

func (s *server) editStudent(ctx *gin.Context) {
	var student demo.Student
	if err := ctx.BindJSON(&student); err != nil {
		s.sendError(ctx, err)
		return
	}

	if err := s.service.EditStudent(&student); err != nil {
		return
	}
	s.send(ctx, student)
}

func (s *server) deleteStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
