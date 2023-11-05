package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"recruitment-task-1/internal/config"
	"time"

	"recruitment-task-1/internal/routes"
)

func main() {
	cfg := config.Load()

	logLevel, err := logger.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	logWriter, err := createLogWriter(cfg.LogFile)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}
	log.Level(logLevel)
	log.Output(logWriter)

	if cfg.ReleaseMode {
		gin.SetMode("release")
	}

	r := gin.New()
	r.Use(ginZerologMiddleware())
	if err := routes.Routes(r, cfg); err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	log.Info().Str("Port", cfg.Port).Msg("Server starting")
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func createLogWriter(filePath string) (io.Writer, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return io.MultiWriter(file, os.Stdout), nil
}

func ginZerologMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)

		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", status).
			Dur("duration", duration).
			Str("ip", c.ClientIP()).
			Msg("request details")
	}
}
