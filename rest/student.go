package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) getStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	student, err := s.service.GetStudent(id)
	if err != nil {
		s.sendErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	s.sendSuccessResponse(ctx, student)
}

func (s *server) addStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *server) editStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *server) deleteStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
