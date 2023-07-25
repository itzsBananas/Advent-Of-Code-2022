package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

type compartment [52]bool

func GetCompartment(s string) compartment {
	var comp compartment

	lowerBase := int('a')
	upperBase := int('A')
	getIndex := func(ch rune) int {
		switch {
		case unicode.IsLower(ch):
			return int(ch) - lowerBase
		case unicode.IsUpper(ch):
			return int(ch) - upperBase + 26
		default:
			log.Fatalf("Invalid char: %v", ch)
		}
		return -1
	}

	for _, ch := range s {
		comp[getIndex(ch)] = true
	}

	return comp
}

func (c compartment) GetCommon(d compartment) compartment {
	var a compartment
	for i := range c {
		a[i] = c[i] && d[i]
	}
	return a
}

func (c compartment) GetScore() int {
	score := 0
	for i, isExist := range c {
		if isExist {
			score += (i + 1)
		}
	}
	return score
}

func (c compartment) String() string {
	var buf bytes.Buffer

	lowerBase := int('a')
	upperBase := int('A')
	getChar := func(i int) rune {
		if i < 0 {
			log.Fatalf("Rune index cannot be less than 0: %d", i)
		} else if i < 26 {
			return rune(i + lowerBase)
		} else if i < 52 {
			return rune(i - 26 + upperBase)
		} else {
			log.Fatalf("Rune index cannot be greater than 52: %d", i)
		}
		return -1
	}
	for index, isExist := range c {
		if isExist {
			buf.WriteRune(getChar(index))
		}
	}
	return buf.String()
}

func (c compartment) size() int {
	size := 0
	for _, isExist := range c {
		if isExist {
			size += 1
		}
	}
	return size
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Invalid number of arg's: %d", len(os.Args))
	}

	fileName := os.Args[1]
	part1(fileName)
	part2(fileName)
}

func part1(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening found from args: %s", err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	totalScore := 0
	for reader.Scan() {
		text := reader.Text()
		first, second := text[:len(text)/2], text[len(text)/2:]
		comp1, comp2 := GetCompartment(first), GetCompartment(second)
		// fmt.Printf("Raw: %s\t%s\n", first, second)
		shared := comp1.GetCommon(comp2)
		score := shared.GetScore()
		// fmt.Printf("Pro: %s \t %s = %s (%d)\n\n", comp1, comp2, shared, score)
		totalScore += score
	}

	fmt.Printf("Total Score (part 1): %d\n", totalScore)
}

func part2(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening found from args: %s", err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	totalScore := 0

	var comp [3]compartment
	i := 0
	for reader.Scan() {
		text := reader.Text()
		comp[i] = GetCompartment(text)
		if i == 2 {
			shared := comp[0].GetCommon(comp[1].GetCommon(comp[2]))
			if shared.size() != 1 {
				log.Fatalf("More than one common item among the three elves: %s", shared)
			}
			// fmt.Printf("%s\t%s\t%s = %s\n", comp[0], comp[1], comp[2], shared)
			totalScore += shared.GetScore()
		}
		i = (i + 1) % 3
	}

	fmt.Printf("Total Score (part 2): %d\n", totalScore)
}
