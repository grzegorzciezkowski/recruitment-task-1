package routes

import (
	"github.com/gin-gonic/gin"
	"recruitment-task-1/internal/config"
	"recruitment-task-1/internal/handlers"
)

func Routes(r *gin.Engine, cfg *config.Config) error {
	numbersHandlers, err := handlers.NewNumbersHandlers(cfg.InputFile)
	if err != nil {
		return err
	}

	r.GET("/api/v1/healthcheck", handlers.Healthcheck)
	r.GET("/api/v1/numbers/:value", numbersHandlers.FindIndex)

	return nil
}
