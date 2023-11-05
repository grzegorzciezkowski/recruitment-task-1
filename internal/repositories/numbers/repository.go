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

func (r *Repository) FindIndex(number int) (int, int, bool) {
	index, ok := r.indexMap[number]
	if ok {
		return index, number, true
	}

	minNumber := number * 90 / 100
	maxNumber := number * 110 / 100

	return r.findLevel10(minNumber, maxNumber, r.numbers)
}

func (r *Repository) findLevel10(min, max int, slice []int) (int, int, bool) {
	if len(slice) == 1 {
		if isBetween(slice[0], min, max) {
			return r.indexMap[slice[0]], slice[0], true
		} else {
			return 0, -1, false
		}
	}

	if !isOverlappingRanges(slice[0], slice[len(slice)-1], min, max) {
		return 0, -1, false
	}

	index, number, ok := r.findLevel10(min, max, slice[:len(slice)/2])
	if ok {
		return index, number, true
	}

	return r.findLevel10(min, max, slice[len(slice)/2:])
}

func isOverlappingRanges(start1, end1, start2, end2 int) bool {
	return start1 <= end2 && end1 >= start2
}

func isBetween(num, min, max int) bool {
	return num >= min && num <= max
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
