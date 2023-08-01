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
	if len(os.Args) != 3 {
		log.Fatalf("Invalid # of arg's: %d", len(os.Args))
	}

	// ignore errors
	knotCount, _ := strconv.Atoi(os.Args[2])

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	b := NewBridge(knotCount)

	for scanner.Scan() {
		line := scanner.Text()
		move := strings.Split(line, " ")
		if len(move) != 2 {
			log.Fatalf("Found incorrect # of arg's (expected 2; got %d)", len(move))
		}

		spaces, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatalf("Could not parse line with int arg; %s of %s", move[1], line)
		}
		switch move[0] {
		case "D":
			// fmt.Printf("down %d\n", spaces)
			b.MoveVertical(-spaces)
		case "U":
			// fmt.Printf("up %d\n", spaces)
			b.MoveVertical(spaces)
		case "R":
			// fmt.Printf("right %d\n", spaces)
			b.MoveHorizontal(spaces)
		case "L":
			// fmt.Printf("left %d\n", spaces)
			b.MoveHorizontal(-spaces)
		default:
			log.Fatalf("Could not parse line with cmd arg; %s of %s", move[0], line)
		}
		// fmt.Println(b)
		// fmt.Println(b.knots)
	}
	fmt.Printf("There are %d positions the tail visited at least once.\n", b.Count())
	// fmt.Println(b)
}
