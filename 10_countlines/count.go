package count

import (
	"bufio"
	"os"
)

// ContLines count lines of content in the file
func CountLines(path string) int {
	var lines int
	// TODO: count lines in file
	dat, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		lines = lines + 1
	}
	return lines
}
