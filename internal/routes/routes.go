package routes

import (
	"github.com/gin-gonic/gin"
	"recruitment-task-1/internal/handlers"
)

func Routes(r *gin.Engine) {
	r.GET("/api/v1/healthcheck", handlers.Healthcheck)
}
