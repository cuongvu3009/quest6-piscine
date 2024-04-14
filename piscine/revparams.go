// Write a program that prints the arguments received in the command line in reverse order.

// Example of output :

// $ go run . choumi is the best cat
// cat
// best
// the
// is
// choumi
// $

// Functions allowed: github.com/01-edu/z01.PrintRune, os.Args,

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Iterate through the arguments received from the command line
	for _, arg := range os.Args[1:] {
		// Print each argument followed by a newline
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
