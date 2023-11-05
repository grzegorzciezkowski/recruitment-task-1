package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	correctFile   = "C:\\Users\\ledyb\\MyApps\\RecruitmentTask1\\recruitment-task-1\\workdir\\input.txt"
	incorrectFile = "C:\\Users\\ledyb\\MyApps\\RecruitmentTask1\\recruitment-task-1\\input_incorrect_number.txt"
)

func TestNewRepository(t *testing.T) {
	testCases := []struct {
		name              string
		filePath          string
		shouldReturnError bool
	}{
		{
			name:     "when file exists",
			filePath: correctFile,
		},
		{
			name:              "when file does not exist",
			filePath:          "fileWhatNotExists",
			shouldReturnError: true,
		},
		{
			name:              "when file has letters in one line",
			filePath:          incorrectFile,
			shouldReturnError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repository, err := NewRepository(tc.filePath)

			if tc.shouldReturnError {
				assert.Error(t, err)
				assert.Equal(t, (*Repository)(nil), repository)
				return
			}

			assert.NoError(t, err)
			assert.NotEqual(t, nil, repository)
		})
	}
}

func TestRepository_FindIndex(t *testing.T) {
	testCases := []struct {
		name           string
		testFile       string
		inputNumber    int
		expectedIndex  int
		expectedNumber int
		numberFound    bool
	}{
		{
			name:           "when number exists",
			testFile:       correctFile,
			inputNumber:    1400,
			expectedIndex:  14,
			expectedNumber: 1400,
			numberFound:    true,
		},
		{
			name:           "when number exists in level 10",
			testFile:       correctFile,
			inputNumber:    1150,
			expectedIndex:  11,
			expectedNumber: 1100,
			numberFound:    true,
		},
		{
			name:           "when number does not exists",
			testFile:       correctFile,
			inputNumber:    115242240,
			expectedIndex:  0,
			expectedNumber: -1,
			numberFound:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repository, err := NewRepository(correctFile)
			assert.NoError(t, err)

			currentIndex, currentNumber, ok := repository.FindIndex(tc.inputNumber)

			assert.Equal(t, tc.expectedIndex, currentIndex)
			assert.Equal(t, tc.expectedNumber, currentNumber)
			assert.Equal(t, tc.numberFound, ok)
		})
	}
}
