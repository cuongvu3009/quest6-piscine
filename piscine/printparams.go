// Write a program that prints the arguments received in the command line.

// Example of output :

// $ go run . choumi is the best cat
// choumi
// is
// the
// best
// cat
// $

// functions allowed: github.com/01-edu/z01.PrintRune, os.Args, --allow-builtin

package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintParams() {
	// Iterate through the arguments received from the command line
	for _, arg := range os.Args[1:] {
		// Print each argument followed by a newline
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
