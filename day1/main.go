package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type top_three [3]int

func (t *top_three) add(a int) {
	i := 0
	for ; i < 3; i++ {
		if a > t[i] {
			break
		}
	}
	for ; i < 3; i++ {
		temp := t[i]
		t[i] = a
		a = temp
	}
}

func (t top_three) String() string {
	return fmt.Sprintf("%d\t%d\t%d", t[0], t[1], t[2])
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

	top := new(top_three)
	total := 0
	line_num := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			top.add(total)
			total = 0
		} else {
			curr, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += curr
		}
		line_num += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	top.add(total)

	fmt.Printf("Max: %s\n", top)
}
