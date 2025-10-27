package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	allFlag := flag.Bool("a", false, "list all files including hidden files")
	longFlag := flag.Bool("l", false, "use a long listing format")

	flag.Parse()

	dirs := flag.Args()
	
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	// fmt.Printf("Flags: -a: %t -l: %t and Directories: %v\n", *allFlag, *longFlag, dirs)

	for _, dir := range dirs {
		fmt.Printf("Listing directory: %s\n", dir)
		
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ls: cannot access '%s': %v\n", dir, err)
			continue
		}

		// fmt.Printf("Number of Directories: %d\n", len(entries))

		for _, entry := range entries {
			name := entry.Name()

			if !*allFlag && name[0] == '.' {
				continue
			}

			if *longFlag {
				info, err := entry.Info()
				if err != nil {
					fmt.Fprintf(os.Stderr, "ls: cannot access info for '%s': %v\n", name, err)
					continue
				}

				mode := info.Mode()
				size := info.Size()
				modTime := info.ModTime()
				fmt.Printf("%s %10d %s %s\n", mode, size, modTime.Format("Jan 02 15:04"), name)
			} else {
				fmt.Println(name)
			}
		}
	}
}
