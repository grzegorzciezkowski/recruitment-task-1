package numbers

import (
	"bufio"
	"os"
	"strconv"
)

type Repository struct {
	filePath string
	numbers  []int
	indexMap map[int]int
}

// NewRepository initializes a new Repository struct by loading data from a specified file.
func NewRepository(filePath string) (*Repository, error) {
	r := &Repository{
		filePath: filePath,
		numbers:  []int{},
		indexMap: make(map[int]int),
	}

	if err := r.loadFile(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Repository) loadFile() error {
	file, err := os.Open(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		r.numbers = append(r.numbers, number)
		r.indexMap[number] = index
		index++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
