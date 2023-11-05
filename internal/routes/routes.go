package routes

import (
	"github.com/gin-gonic/gin"
	"recruitment-task-1/internal/handlers"
)

func Routes(r *gin.Engine) error {
	numbersHandlers, err := handlers.NewNumbersHandlers()
	if err != nil {
		return err
	}

	r.GET("/api/v1/healthcheck", handlers.Healthcheck)
	r.GET("/api/v1/numbers/:value", numbersHandlers.FindIndex)

	return nil
}
