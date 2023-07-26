package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type stack []rune

func (s stack) String() string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	for _, ch := range s {
		buf.WriteRune(ch)
		buf.WriteRune(',')
	}
	buf.WriteRune(']')
	return buf.String()
}

func (s stack) Reverse() stack {
	rev := len(s) - 1
	for i := range s[:len(s)/2] {
		s[i], s[rev] = s[rev], s[i]
		rev -= 1
	}
	return s
}

func (s stack) Peek() rune {
	return s[len(s)-1]
}

func (s stack) Read(a int) stack {
	if a < 0 {
		log.Fatal("Pop value (a) cannot be less than 0")
	}
	return s[len(s)-a:]
}

func (s stack) Pop(a int) stack {
	if a < 0 {
		log.Fatal("Pop value (a) cannot be less than 0")
	}
	return s[:len(s)-a]
}

func (s stack) Push(n stack) stack {
	return append(s, n...)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Invalid # of arg's: %d", len(os.Args))
	}
	fmt.Println("Part 1: Reverse Order")
	part1(os.Args[1])
	fmt.Println("\nPart 2: Keep Order")
	part2(os.Args[1])
}

func part1(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true

	var groupCount int
	var group []stack
	const stepSize = 4
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			// find # of groups based on number of char's (not reliable but servicable)
			groupCount = (len(line) + 1) / stepSize
			for i := 0; i < groupCount; i++ {
				group = append(group, make(stack, 0))
			}
			firstLine = false
		}
		groupIndex := 0
		for i := 1; i < len(line); i += stepSize {
			if groupIndex >= len(group) {
				break
			}
			r := rune(line[i])
			if unicode.IsUpper(r) {
				group[groupIndex] = append(group[groupIndex], r)
			}
			groupIndex += 1
		}
		if len(line) == 0 {
			break
		}
	}

	fmt.Println("Starting stacks of crates:")
	for i, s := range group {
		group[i] = s.Reverse()
		fmt.Printf("%d: %s\n", i, s)
	}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if len(words) != 6 {
			log.Fatalf("Line cannot be parsed due to # of spaces (%d): %s", len(words)-1, line)
		}
		quant, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		from -= 1
		to, _ := strconv.Atoi(words[5])
		to -= 1
		// fmt.Printf("%s -%d-> %s\n", group[from], quant, group[to])
		popped := group[from].Read(quant).Reverse()
		group[to] = group[to].Push(popped)
		group[from] = group[from].Pop(quant)
		// fmt.Printf("%s -%d-> %s\n", group[from], quant, popped)
		// fmt.Println()
	}

	fmt.Print("Top of the stack afterwards: ")
	for _, ch := range group {
		fmt.Printf("%s", string(ch.Peek()))
	}
	fmt.Println()
}

func part2(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true

	var groupCount int
	var group []stack
	const stepSize = 4
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			// find # of groups based on number of char's (not reliable but servicable)
			groupCount = (len(line) + 1) / stepSize
			for i := 0; i < groupCount; i++ {
				group = append(group, make(stack, 0))
			}
			firstLine = false
		}
		groupIndex := 0
		for i := 1; i < len(line); i += stepSize {
			if groupIndex >= len(group) {
				break
			}
			r := rune(line[i])
			if unicode.IsUpper(r) {
				group[groupIndex] = append(group[groupIndex], r)
			}
			groupIndex += 1
		}
		if len(line) == 0 {
			break
		}
	}

	fmt.Println("Starting stacks of crates:")
	for i, s := range group {
		group[i] = s.Reverse()
		fmt.Printf("%d: %s\n", i, s)
	}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if len(words) != 6 {
			log.Fatalf("Line cannot be parsed due to # of spaces (%d): %s", len(words)-1, line)
		}
		quant, _ := strconv.Atoi(words[1])
		from, _ := strconv.Atoi(words[3])
		from -= 1
		to, _ := strconv.Atoi(words[5])
		to -= 1
		// fmt.Printf("%s -%d-> %s\n", group[from], quant, group[to])
		popped := group[from].Read(quant)
		group[to] = group[to].Push(popped)
		group[from] = group[from].Pop(quant)
		// fmt.Printf("%s -%d-> %s\n", group[from], quant, popped)
		// fmt.Println()
	}

	fmt.Print("Top of the stack afterwards: ")
	for _, ch := range group {
		fmt.Printf("%s", string(ch.Peek()))
	}
	fmt.Println()
}
