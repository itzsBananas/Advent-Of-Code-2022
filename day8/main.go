package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type grid [][]int

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

	forest := make(grid, 0)
	visible := make(grid, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		forest = append(forest, make([]int, len(line)))
		visible = append(visible, make([]int, len(line)))
		if len(line) == 0 {
			continue
		}

		for j, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatalf("Failed parsing line (%s) with char (%c) as int", line, ch)
			}
			forest[i][j] = num
			visible[i][j] = 0
		}

		i += 1
	}
	part1(forest, visible)
	part2(forest)
}

func part1(forest grid, visible grid) {
	traversals := [4]<-chan coord{TopLeft(forest), BottomLeft(forest), TopRight(forest), BottomRight(forest)}
	for _, t := range traversals {
		maxHeight := 0
		for c := range t {
			x, y := c.x, c.y
			val := c.val
			newIteration := c.newIteration
			if newIteration {
				visible[x][y] = 1
				maxHeight = val
				continue
			}

			if val > maxHeight {
				visible[x][y] = 1
				maxHeight = val
			}
		}
		// fmt.Println(visible)
	}

	t := TopLeft(forest)
	visibleTrees := 0
	for c := range t {
		x, y := c.x, c.y
		if visible[x][y] != 0 {
			visibleTrees += 1
		}
	}
	fmt.Printf("# of visible trees: %d\n", visibleTrees)
}

func part2(forest grid) {
	scores := make(grid, 0)
	highestScore := 0

	for i, row := range forest {
		scores = append(scores, make([]int, len(row)))
		for j, tree := range row {
			// left
			score := 1
			x := i
			for ; x > 0; x-- {
				if tree <= forest[x-1][j] {
					x -= 1
					break
				}
			}
			score *= i - x
			// right
			x = i
			for ; x < len(forest)-1; x++ {
				if tree <= forest[x+1][j] {
					x += 1
					break
				}
			}
			score *= x - i
			// up
			y := j
			for ; y > 0; y-- {
				if tree <= forest[i][y-1] {
					y -= 1
					break
				}
			}
			score *= j - y
			// down
			y = j
			for ; y < len(forest[i])-1; y++ {
				if tree <= forest[i][y+1] {
					y += 1
					break
				}
			}
			score *= y - j
			scores[i][j] = score
			if score > highestScore {
				highestScore = score
			}
		}
	}
	fmt.Printf("The highest scenic score possible for any tree is %d\n", highestScore)
}
