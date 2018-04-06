package main

import "fmt"

func main() {
	i := []int{0, 0, 0, 0}
	f := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	s := []string{"dddd", "dddd", "dsdd", "ssssss"}
	// Declare a variable called i which is a slice of 5 int.
	// Declare a variable called f which is a slice of 9 float64.
	// Declare a variable called s which is a slice of 4 string.

	fmt.Println(len(i), len(f), len(s))
}
