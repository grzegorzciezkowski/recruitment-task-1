package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRoutes(t *testing.T) {
	expectedRoutesInfo := gin.RoutesInfo{
		gin.RouteInfo{
			Method:  "GET",
			Path:    "/api/v1/healthcheck",
			Handler: "recruitment-task-1/internal/handlers.Healthcheck",
		},
	}

	r := gin.Default()
	Routes(r)
	currentRoutesInfo := r.Routes()

	for i, currentRoute := range currentRoutesInfo {
		assert.Equal(t, expectedRoutesInfo[i].Method, currentRoute.Method)
		assert.Equal(t, expectedRoutesInfo[i].Path, currentRoute.Path)
		assert.Equal(t, expectedRoutesInfo[i].Handler, currentRoute.Handler)
	}
}
