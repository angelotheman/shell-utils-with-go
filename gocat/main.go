package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	nFlag := flag.Bool("n", false, "number all output lines")
	EFlag := flag.Bool("E", false, "display $ at end of each line")

	flag.Parse()

	files := flag.Args()
	
	if len(files) == 0 {
		fmt.Println("Usage: go run cat.go [-n] [-E] <filename>")
		return
	}

	// Looping over each file provided as argument

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("cat: %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 1

		for scanner.Scan() {
			line := scanner.Text()

			if *EFlag {
				line += "$"
			}

			if *nFlag {
				fmt.Printf("%6d  %s\n", lineNum, line)
				lineNum++
			} else {
				fmt.Println(line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file %s: %v\n", filename, err)
		}
	}
}
