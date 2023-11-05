package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"recruitment-task-1/internal/config"

	"testing"
)

const (
	correctFile   = "../../testdata/input.txt"
	incorrectFile = "../../testdata/input_incorrect_number.txt"
)

func TestRoutes(t *testing.T) {
	expectedRoutesInfo := gin.RoutesInfo{
		gin.RouteInfo{
			Method:  "GET",
			Path:    "/api/v1/healthcheck",
			Handler: "recruitment-task-1/internal/handlers.Healthcheck",
		},
		gin.RouteInfo{
			Method:  "GET",
			Path:    "/api/v1/numbers/:value",
			Handler: "recruitment-task-1/internal/handlers.(*NumbersHandlers).FindIndex-fm",
		},
	}

	testCases := []struct {
		name              string
		cfg               *config.Config
		shouldReturnError bool
	}{
		{
			name: "with correct config file",
			cfg: &config.Config{
				InputFile: correctFile,
			},
			shouldReturnError: false,
		},
		{
			name: "with incorrect config file",
			cfg: &config.Config{
				InputFile: incorrectFile,
			},
			shouldReturnError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.Default()

			err := Routes(r, tc.cfg)
			if tc.shouldReturnError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			currentRoutesInfo := r.Routes()
			for i, currentRoute := range currentRoutesInfo {
				assert.Equal(t, expectedRoutesInfo[i].Method, currentRoute.Method)
				assert.Equal(t, expectedRoutesInfo[i].Path, currentRoute.Path)
				assert.Equal(t, expectedRoutesInfo[i].Handler, currentRoute.Handler)
			}
		})
	}
}
