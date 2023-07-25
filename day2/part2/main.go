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

type outcome struct {
	ch rune
}

func (o outcome) getOutcomeBaseScore(m move) int {
	var delta int
	switch o.ch {
	case 'X':
		delta = -1
	case 'Y':
		delta = 0
	case 'Z':
		delta = 1
	default:
		log.Fatalf("Illegal outcome: %q", o)
	}
	score := m.moveIndex + delta
	if score < 0 {
		score += numberOfMoves
	} else if score >= numberOfMoves {
		score -= numberOfMoves
	}
	return move{score}.getBaseScore()
}

func (o outcome) getOutcomeBonusScore() int {
	switch o.ch {
	case 'X':
		return 0
	case 'Y':
		return 3
	case 'Z':
		return 6
	default:
		log.Fatalf("Illegal outcome: %q", o)
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
		opp, outcome := getMove(runes[0]), outcome{runes[2]}
		totalScore += outcome.getOutcomeBaseScore(opp)
		totalScore += outcome.getOutcomeBonusScore()
		// fmt.Printf("%q with outcome %q: %d\n", runes[0], runes[2], outcome.getOutcomeBaseScore(opp))
	}
	fmt.Printf("Total Score: %d\n", totalScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func test() {
	moves0 := []rune{'A', 'B', 'C'}
	outcomes := []rune{'X', 'Y', 'Z'}
	for _, m0 := range moves0 {
		for _, out := range outcomes {
			move0 := getMove(m0)
			o := outcome{out}
			fmt.Printf("%q with outcome %q: %d\n", m0, o, o.getOutcomeBaseScore(move0))
		}
	}
}
