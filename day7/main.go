package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Invalid # of arg's: %d", len(os.Args))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fs := CreateFileSystem()
	isListState := false
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		// is a command
		if line[0] == '$' {
			if len(line) < 4 {
				log.Fatalf("Command too short: %s", line)
			}
			cmd := line[2:4]
			switch cmd {
			case "ls":
				isListState = true
			case "cd":
				isListState = false
				fs.ChangeDirectory(line[5:])
			}
			continue
		}

		if !isListState {
			log.Fatalf("File or dir was listed before a 'ls' command was issued: %s", line)
		}
		// if not list state, panic
		sep := strings.Split(line, " ")
		if len(sep) != 2 {
			log.Fatalf("ls command found file or directory w/ wrong formatting: %s", sep)
		}

		if sep[0] == "dir" {
			dir := CreateDir(sep[1])
			fs.AddDir(&dir)
		} else {
			// add file
			size, err := strconv.Atoi(sep[0])
			if err != nil {
				log.Fatalf("Listed file cannot parse filesize: %s", err)
			}
			name := sep[1]
			file := CreateFile(name, size)
			fs.AddFile(&file)
		}
	}

	// Part A
	root := fs.rootDir
	totalSize := 0
	const bound = 100000
	check := func(size int) {
		if size < bound {
			totalSize += size
		}
	}
	usedSpace := root.TotalSize(check)
	fmt.Printf("Part A: Total Size of those <100000 directory: %d\n", totalSize)

	// Part B
	root = fs.rootDir
	largestSize := int(^uint(0) >> 1)

	const totalSpace = 70000000
	const requiredUnusedSpace = 30000000

	remaining := requiredUnusedSpace - (totalSpace - usedSpace)
	if remaining <= 0 {
		fmt.Println("Part B: No need to increase unused space")
	}
	check = func(size int) {
		if size < largestSize && size > remaining {
			largestSize = size
		}
	}
	root.TotalSize(check)
	fmt.Printf("Part B: The size of the minimum costing, viable directory to free up space is %d\n", largestSize)
}
