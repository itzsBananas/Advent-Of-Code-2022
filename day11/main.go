package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const outerVerbose = false

type operate func(old int) int

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Expected %d cmd-line arg's but recieved %d cmd-line arg's instead", 2, len(os.Args))
	}
	paragraphs := newGenerator(os.Args[1])
	numOfRounds, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Unable to parse numOfRounds arg: %s\n", err)
	}

	monkeys := initMonkeys(paragraphs)
	for i := 1; i < numOfRounds+1; i++ {
		monkeys.PlayRound()
		if outerVerbose {
			fmt.Printf("After round %d, the monkeys are holding items with these worry levels:\n", i)
			fmt.Println(monkeys)
		}
	}

	opCount := monkeys.InspectCount()
	if outerVerbose {
		for i, c := range opCount {
			fmt.Printf("Monkey %d inspected items %d times.\n", i, c)
		}
	}
	sort.Ints(opCount)

	lastIndex := len(opCount) - 1
	fmt.Printf("The level of monkey business after %d rounds of stuff-slinging simian shenanigans was %d\n",
		numOfRounds,
		opCount[lastIndex]*opCount[lastIndex-1])
}

const paragraphSize = 6

func newGenerator(fileName string) <-chan [paragraphSize]string {
	ch := make(chan [paragraphSize]string)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file named %s: %s", fileName, err)
	}

	scanner := bufio.NewScanner(file)
	go func() {
		var output [paragraphSize]string
		var currLine int

		for scanner.Scan() {
			if currLine == paragraphSize {
				currLine = 0
				continue
			}
			output[currLine] = scanner.Text()

			if currLine == paragraphSize-1 {
				ch <- output
			}
			currLine += 1
		}

		defer file.Close()
		close(ch)
	}()

	return ch
}

func initMonkeys(ch <-chan [paragraphSize]string) monkeyGame {
	monkeys := make([]*monkey, 0)
	minDivisor := 1
	for para := range ch {
		items := processItemLine(trimLine(para[1]))
		op := processOpLine(trimLine(para[2]))
		div := processDivisorLine(trimLine(para[3]))
		trueMonkey := processNextMonkeyLine(trimLine(para[4]))
		falseMonkey := processNextMonkeyLine(trimLine(para[5]))

		monkeys = append(monkeys, NewMonkey(items, op, div, trueMonkey, falseMonkey))
		minDivisor *= div
	}
	// part a
	// return monkeyGame{monkeys: monkeys, divisor: 3}
	// part b
	return monkeyGame{monkeys: monkeys, minCommonDivisor: minDivisor}
}

func trimLine(str string) string {
	cmd := strings.Split(str, ": ")
	if len(cmd) != 2 {
		log.Fatalf("Expected two string entries sep by ': ' in starting items row: recieved %s", cmd)
	}
	return cmd[1]
}

func processItemLine(str string) []int {
	itemStrings := strings.Split(str, ", ")
	items := make([]int, len(itemStrings))
	for i := range itemStrings {
		val, err := strconv.Atoi(itemStrings[i])
		if err != nil {
			log.Fatalf("Unable to parse monkey items: %s", err)
		}
		items[i] = int(val)
	}
	return items
}

func processOpLine(str string) operate {
	cmd := strings.Split(str, "= ")
	if len(cmd) != 2 {
		log.Fatalf("Expected two string entries sep by '= ' in op row: recieved %s", cmd)
	}
	eq := cmd[1]
	ops := strings.Split(eq, " ")
	if len(ops) != 3 {
		log.Fatalf("Expected three string entries sep by ' ' in op row: recieved %s", ops)
	}

	operand := func(str string) operate {
		switch str {
		case "old":
			return func(old int) int { return old }
		default:
			operand, err := strconv.Atoi(str)
			if err != nil {
				log.Fatalf("Unable to parse operand for op: (%s)", str)
			}
			return func(_ int) int { return int(operand) }
		}
	}

	operator := func(str string, old int) func(op1, op2 operate) int {
		switch str {
		case "*":
			return func(op1, op2 operate) int { return op1(old) * op2(old) }
		case "+":
			return func(op1, op2 operate) int { return op1(old) + op2(old) }
		default:
			log.Fatalf("Unable to parse operator for op: %s", str)
			return func(_, _ operate) int { return 0 }
		}
	}

	return func(old int) int { f := operator(ops[1], old); return f(operand(ops[0]), operand(ops[2])) }
}

func processDivisorLine(str string) int {
	cmd := strings.Split(str, " ")
	if len(cmd) != 3 {
		log.Fatalf("Expected three string entries sep by ' ' in processDivisor row: recieved %s", cmd)
	}
	val, err := strconv.Atoi(cmd[2])
	if err != nil {
		log.Fatalf("Unable to parse divisor: %s", err)
	}

	return int(val)
}

func processNextMonkeyLine(str string) int {
	cmd := strings.Split(str, " ")
	if len(cmd) != 4 {
		log.Fatalf("Expected four string entries sep by ' ' in monkey row: recieved %s", cmd)
	}
	val, err := strconv.Atoi(cmd[3])
	if err != nil {
		log.Fatalf("Unable to parse monkey index: %s", err)
	}

	return val
}
