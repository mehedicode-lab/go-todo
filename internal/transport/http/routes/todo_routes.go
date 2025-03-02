package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/internal/transport/http/handlers"
)

func SetupRoutes(handler *handlers.TodoHandler) *gin.Engine {
	r := gin.Default()

	todo := r.Group("/todos")
	{
		todo.POST("/", handler.Create)
		todo.GET("/", handler.GetAll)
		todo.GET("/:id", handler.GetByID)
		todo.PUT("/:id", handler.Update)
		todo.DELETE("/:id", handler.Delete)
	}

	return r
}
