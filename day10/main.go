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

	gen := readFile(os.Args[0])
	cpu := New()
	for line := range gen {
		args := strings.Split(line, " ")
		if len(args) > 2 || len(args) == 0 {
			log.Fatalf("Invalid # of arg's: %d", len(args))
		}

		switch args[0] {
		case "addx":
			if len(args) != 2 {
				log.Fatalf("Invalid # of arg's for addx: %d (expected: 2)", len(args))
			}

			val, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("Could not parse int arg for addx: %s", err)
			}
			cpu.Addx(val)
		case "noop":
			if len(args) != 1 {
				log.Fatalf("Invalid # of arg's for noop: %d (expected: 1)", len(args))
			}
			cpu.Noop()

		}
	}
	fmt.Printf("Part A: %d\n", cpu.Sum())
}

func readFile(name string) <-chan string {
	ch := make(chan string)

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Something wrong with readFile: %s", err)
		}
		defer file.Close()
		close(ch)
	}()
	return ch
}
