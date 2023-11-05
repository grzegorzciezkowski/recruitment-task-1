package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"recruitment-task-1/internal/repositories/numbers"
	"strconv"
)

type NumbersHandlers struct {
	numbersRepository *numbers.Repository
}

type MessageResponse struct {
	Message string `json:"message"`
}

type FindIndexResponse struct {
	Index  int `json:"index"`
	Number int `json:"value"`
}

func NewNumbersHandlers() (*NumbersHandlers, error) {
	repository, err := numbers.NewRepository("C:\\Users\\ledyb\\MyApps\\RecruitmentTask1\\recruitment-task-1\\workdir\\input.txt")
	if err != nil {
		return nil, err
	}

	return &NumbersHandlers{numbersRepository: repository}, nil
}

func (h *NumbersHandlers) FindIndex(g *gin.Context) {
	value, err := strconv.Atoi(g.Param("value"))
	if err != nil {
		log.Printf(err.Error())
		return
	}

	index, number, ok := h.numbersRepository.FindIndex(value)
	if !ok {
		g.JSON(http.StatusNotFound, MessageResponse{
			Message: fmt.Sprintf("Number %d has not been found.", value),
		})
		return
	}

	g.JSON(http.StatusOK, FindIndexResponse{
		Index:  index,
		Number: number,
	})
}
