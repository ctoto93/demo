package rest

import "github.com/gin-gonic/gin"

func (s *server) getStudent(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
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
