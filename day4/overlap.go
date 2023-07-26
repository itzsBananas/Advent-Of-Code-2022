package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start, end int
}

func (s section) FullyOverlap(o section) bool {
	if s.start > o.start {
		return o.FullyOverlap(s)
	}
	return s.start == o.start || s.end >= o.end
}

func (s section) Overlap(o section) bool {
	if s.start > o.start {
		return o.Overlap(s)
	}
	return s.end >= o.start
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Invalid # of arg's: %d", len(os.Args))
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	Ocount := 0
	FOcount := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		if len(split) != 2 {
			log.Fatalf("Line input [%s] has invalid # of sections: %d", line, len(split))
		}
		var pairs [2]section

		for i, sec := range split {
			endpoints := strings.Split(sec, "-")
			if len(endpoints) != 2 {
				log.Fatalf("Section [%s] has invalid # of sections: %d", sec, len(endpoints))
			}
			start, err := strconv.Atoi(endpoints[0])
			if err != nil {
				log.Fatalf("Start point of section [%s] misread: %s", sec, err)
			}
			end, err := strconv.Atoi(endpoints[1])
			if err != nil {
				log.Fatalf("End point of section [%s] misread: %s", sec, err)
			}
			if start > end {
				log.Fatalf("Start point cannot be greater than end point: %d / %d", start, end)
			}

			pairs[i] = section{start, end}
		}

		fo := pairs[0].FullyOverlap(pairs[1])
		o := pairs[0].Overlap(pairs[1])
		if fo {
			FOcount += 1
		}
		if o {
			Ocount += 1
		}
		// fmt.Printf("%v = %t, %t\n", pairs, fo, o)
	}
	fmt.Printf("Total full overlaps: %d\n", FOcount)
	fmt.Printf("Total overlaps: %d\n", Ocount)
}
