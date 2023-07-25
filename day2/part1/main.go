package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type move struct {
	moveIndex int
}

const numberOfMoves = 3

func getMove(char rune) move {
	var m move
	switch char {
	case 'A', 'X':
		m.moveIndex = 0
	case 'B', 'Y':
		m.moveIndex = 1
	case 'C', 'Z':
		m.moveIndex = 2
	default:
		log.Fatalf("Invalid move: %q", char)
	}
	return m
}

func (m move) getTotalScore(o move) int {
	return m.getBaseScore() + m.getBonusScore(o)
}

func (m move) getBaseScore() int {
	return m.moveIndex + 1
}

// 0 = tie, 1 = win, 2 = lose
func (m move) getBonusScore(o move) int {
	diff := m.moveIndex - o.moveIndex
	if diff < 0 {
		diff += numberOfMoves
	}
	switch diff {
	case 0:
		return 3
	case 1:
		return 6
	case 2:
		return 0
	}
	return 0
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Invalid # of arg's")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		if len(runes) != 3 || runes[1] != rune(' ') {
			log.Fatalf("Illegal formatting: %s:", line)
		}
		m0, m1 := getMove(runes[0]), getMove(runes[2])
		// fmt.Printf("%q * %q = %d\n", runes[0], runes[2], m1.getTotalScore(m0))
		totalScore += m1.getTotalScore(m0)
	}
	fmt.Printf("Total Score: %d\n", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func test() {
	moves0 := []rune{'A', 'B', 'C'}
	moves1 := []rune{'X', 'Y', 'Z'}
	for _, m0 := range moves0 {
		for _, m1 := range moves1 {
			move0 := getMove(m0)
			move1 := getMove(m1)
			fmt.Printf("%q * %q = %d\n", m0, m1, move1.getTotalScore(move0))
		}
	}
}
