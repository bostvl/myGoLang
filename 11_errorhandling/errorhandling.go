package errorhandling

import (
	"bufio"
	"os"
)

// CountLines count lines in the file
func CountLines(path string) (int, error) {
	var lines int
	var err error
	// TODO: count lines in file
	dat, err := os.Open(path)

	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		lines = lines + 1
	}
	return lines, err
}
