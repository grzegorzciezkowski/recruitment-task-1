package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const statusOK = "OK"

type healthcheckResponse struct {
	Status string `json:"status"`
}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthcheckResponse{Status: statusOK})
}
