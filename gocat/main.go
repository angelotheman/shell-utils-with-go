package main

import (
	"fmt"
	"os"
)

func main() {
	// Print the arguments passed to the program
	args := os.Args

	fmt.Println(args)

	fc, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Println(string(fc))
}
