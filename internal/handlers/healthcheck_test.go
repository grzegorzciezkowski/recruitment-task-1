package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	recorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(recorder)
	req := &http.Request{
		URL: &url.URL{},
	}
	ginContext.Request = req

	Healthcheck(ginContext)

	var currentResponse healthcheckResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &currentResponse)
	assert.NoError(t, err)

	expectedResponse := healthcheckResponse{Status: statusOK}
	assert.Equal(t, expectedResponse, currentResponse)
}
