package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

type recent struct {
	arr  []rune
	size int
}

func NewRecent(cap int) recent {
	return recent{make([]rune, cap), 0}
}

func (r recent) String() string {
	var buf bytes.Buffer
	for i := 0; i < len(r.arr); i++ {
		if r.arr[i] == 0 {
			break
		}
		buf.WriteRune(r.arr[i])
	}
	return buf.String()
}

func (r recent) Has(ru rune) bool {
	for i := 0; i < len(r.arr); i++ {
		if r.arr[i] == ru {
			return true
		}
	}
	return false
}

func (r recent) Add(ru rune) recent {
	for i := 0; i < len(r.arr); i++ {
		if r.arr[i] == 0 {
			r.arr[i] = ru
			r.size += 1
			break
		} else if r.arr[i] == ru {
			break
		}
	}
	return r
}

func (r recent) Clear(ru rune) recent {
	i := 0
	for ; i < len(r.arr); i++ {
		if r.arr[i] == ru {
			r.arr[i] = 0
			i += 1
			break
		}
		r.arr[i] = 0
	}
	gap := i
	r.size = r.size - i
	for ; i < len(r.arr); i++ {
		if r.arr[i] == 0 {
			break
		}
		r.arr[i-gap] = r.arr[i]
		r.arr[i] = 0
	}
	return r
}

func (r *recent) IsFull() bool {
	return len(r.arr) == r.size
}

// args: [fileName] [markerSize]
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Invalid # of arg's: %d", len(os.Args))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	markerSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Error parsing argument 2: %s", err)
	}

	rec := NewRecent(markerSize)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		i += 1

		if l := len(line); l != 1 {
			log.Fatalf("Scan expected string length of 1; got %d", l)
		}
		ru := rune(line[0])
		if rec.Has(ru) {
			rec = rec.Clear(ru)
		}
		rec = rec.Add(ru)
		if rec.IsFull() {
			break
		}
		// fmt.Println(rec)
	}
	if !rec.IsFull() {
		fmt.Println("No marker found :(")
		return
	}
	fmt.Printf("First Marker: %s\n", rec)
	fmt.Printf("Index %d\n", i)
}
