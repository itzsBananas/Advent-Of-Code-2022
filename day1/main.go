package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Invalid # of arg's")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	most, total := 0, 0
	line_num := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if total > most {
				most = total
			}
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

	if total > most {
		most = total
		total = 0
	}

	fmt.Printf("Max: %d\n", most)
}
