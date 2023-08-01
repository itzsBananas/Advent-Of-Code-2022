package main

import "fmt"

type CPU struct {
	currCycle    int
	regVal       int
	sum          int
	nextKeyCycle int
}

func New() CPU {
	return CPU{currCycle: 0, regVal: 1, nextKeyCycle: 40}
}

func (c *CPU) Addx(i int) {
	c.addCycle()
	c.addCycle()
	c.regVal += i
}

func (c *CPU) Noop() {
	c.addCycle()
}

func (c *CPU) addCycle() {
	// fmt.Println(c.currCycle, c.nextKeyCycle)
	c.printCycle()
	c.currCycle += 1
	if c.currCycle >= c.nextKeyCycle {
		c.sum += c.regVal * c.currCycle
		c.currCycle = 0
		// fmt.Printf("At cycle %d: %d\n", c.currCycle, c.regVal)
	}
}

func (c CPU) printCycle() {
	diff := c.currCycle - c.regVal
	if diff < 2 && diff > -2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if c.currCycle+1 == c.nextKeyCycle {
		fmt.Println("")
	}
}

func (c CPU) Sum() int {
	return c.sum
}
