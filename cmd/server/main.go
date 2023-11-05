package main

import (
	"github.com/gin-gonic/gin"

	"recruitment-task-1/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
