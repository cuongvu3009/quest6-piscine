// Write a program that prints the name of the program.

// Example of output :

// student/piscine/printprogramname$ go build main.go
// student/piscine/printprogramnane$ ./main
// main
// student/piscine/printprogramname$ go build
// student/piscine/printprogramname$ ./printprogramname | cat -e
// printprogramname$
// student/piscine/printprogramname$ go build -o Nessy
// student/piscine/printprogramname$ ./Nessy
// Nessy
// student/piscine/printprogramname$

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Exclude the program name from the arguments
	// Print arguments in ASCII order
	for len(args) > 0 {
		minIndex := 0
		// Find the index of the argument with the smallest ASCII character
		for i := 1; i < len(args); i++ {
			if compare(args[i], args[minIndex]) < 0 {
				minIndex = i
			}
		}
		// Print the argument with the smallest ASCII character
		for _, char := range args[minIndex] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
		// Remove the printed argument from the slice
		args = append(args[:minIndex], args[minIndex+1:]...)
	}
}

// compare compares two strings lexicographically based on their ASCII values.
// It returns a negative value if s1 is less than s2, 0 if they are equal,
// and a positive value if s1 is greater than s2.
func compare(s1, s2 string) int {
	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
		i++
	}
	return len(s1) - len(s2)
}
