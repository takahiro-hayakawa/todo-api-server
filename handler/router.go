package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouting(e *gin.Engine, todoHandler TodoHandler) {
	e.GET("/", todoHandler.View())
	e.GET("/:id", todoHandler.Find())
}
