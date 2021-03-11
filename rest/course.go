package rest

import "github.com/gin-gonic/gin"

func (s *server) getCourse(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *server) addCourse(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *server) editCourse(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *server) deleteCourse(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
