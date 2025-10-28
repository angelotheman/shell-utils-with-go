package main

import (
	"flag"
	"fmt"
)

func main () {
	nFlag := flag.Int("n", 10, "number of lines to display from the start")
	flag.Parse()
	
	lines := flag.Args()
	if len(lines) == 0 {
		fmt.Println("No input lines given")
	}

	fmt.Printf("%v\n", nFlag)
}