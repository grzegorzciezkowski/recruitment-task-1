package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepository(t *testing.T) {
	testCases := []struct {
		name              string
		filePath          string
		shouldReturnError bool
	}{
		{
			name:     "when file exists",
			filePath: "C:\\Users\\ledyb\\MyApps\\RecruitmentTask1\\recruitment-task-1\\workdir\\input.txt",
		},
		{
			name:              "when file does not exist",
			filePath:          "fileWhatNotExists",
			shouldReturnError: true,
		},
		{
			name:              "when file has letters in one line",
			filePath:          "C:\\Users\\ledyb\\MyApps\\RecruitmentTask1\\recruitment-task-1\\input_incorrect_number.txt",
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
