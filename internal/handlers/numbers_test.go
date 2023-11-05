package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	correctFile   = "../../testdata/input.txt"
	incorrectFile = "../../testdata/input_incorrect_number.txt"
)

func TestNewNumbersHandlers(t *testing.T) {
	testCases := []struct {
		name              string
		inputFile         string
		shouldReturnError bool
	}{
		{
			name:              "with correct configuration",
			inputFile:         correctFile,
			shouldReturnError: false,
		},
		{
			name:              "with incorrect configuration",
			inputFile:         incorrectFile,
			shouldReturnError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			numbersHandlers, err := NewNumbersHandlers(tc.inputFile)

			if tc.shouldReturnError {
				assert.Error(t, err)
				assert.Nil(t, numbersHandlers)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, numbersHandlers)
		})
	}
}

func TestNumbersHandlers_FindIndex(t *testing.T) {
	testCases := []struct {
		name             string
		value            string
		expectedHttpCode int
		expectedResponse gin.H
	}{
		{
			name:             "when number has been found",
			value:            "1400",
			expectedHttpCode: http.StatusOK,
			expectedResponse: gin.H{"index": float64(14), "value": float64(1400)},
		},
		{
			name:             "when number has not been found",
			value:            "11525424340",
			expectedHttpCode: http.StatusNotFound,
			expectedResponse: gin.H{"message": "Number 11525424340 has not been found."},
		},
		{
			name:             "when value is number",
			value:            "aaaa",
			expectedHttpCode: http.StatusBadRequest,
			expectedResponse: gin.H{"message": "strconv.Atoi: parsing \"aaaa\": invalid syntax"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ginContext, _ := gin.CreateTestContext(recorder)
			req := &http.Request{
				URL: &url.URL{},
			}
			ginContext.Request = req
			ginContext.Params = append(ginContext.Params, gin.Param{
				Key:   "value",
				Value: tc.value,
			})

			numbersHandlers, err := NewNumbersHandlers(correctFile)
			assert.NoError(t, err)

			numbersHandlers.FindIndex(ginContext)
			var currentResponse gin.H
			err = json.Unmarshal(recorder.Body.Bytes(), &currentResponse)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedHttpCode, recorder.Code)
			assert.Equal(t, tc.expectedResponse, currentResponse)
		})
	}
}
