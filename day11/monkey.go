package main

import (
	"bytes"
	"fmt"
)

const innerVerbose = false
const divisor = 3

type test struct {
	divisor     int
	trueMonkey  int
	falseMonkey int
}

type monkey struct {
	items   []int
	op      func(old int) int
	opCount int
	test
}

type monkeyGame struct {
	monkeys []*monkey
	divisor int
}

func (g *monkeyGame) PlayRound() {
	length := len(g.monkeys)
	for i := 0; i < length; i++ {
		if innerVerbose {
			fmt.Printf("Monkey %d:\n", i)
		}
		g.playTurn(i)
	}
}

func (g *monkeyGame) playTurn(i int) {
	slice := g.monkeys
	m := slice[i]
	for _, item := range m.items {
		opLevel := m.op(item)
		reliefLevel := opLevel / divisor
		isDivisble := reliefLevel%m.divisor == 0

		var nextMonkey int
		if isDivisble {
			nextMonkey = m.trueMonkey
			slice[nextMonkey].items = append(slice[nextMonkey].items, reliefLevel)
		} else {
			nextMonkey = m.falseMonkey
			slice[nextMonkey].items = append(slice[nextMonkey].items, reliefLevel)
		}
		if innerVerbose {
			fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
			fmt.Printf("    Worry level is updated to %d.\n", opLevel)
			fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", reliefLevel)
			fmt.Printf("    Current worry level is not divisible by %d.\n", m.divisor)
			fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", reliefLevel, nextMonkey)
		}
	}
	slice[i].opCount += len(m.items)
	slice[i].items = slice[i].items[:0]

	g.monkeys = slice
}

func (g monkeyGame) InspectCount() []int {
	count := make([]int, len(g.monkeys))
	for i, c := range g.monkeys {
		count[i] = c.opCount
	}
	return count
}

func (g monkeyGame) String() string {
	var b bytes.Buffer

	for i, m := range g.monkeys {
		b.WriteString(fmt.Sprintf("Monkey %d: %v\n", i, m.items))
	}

	return b.String()
}

func NewMonkey(items []int, op func(old int) int, divisor int, trueMonkey int, falseMonkey int) *monkey {
	return &monkey{items, op, 0, test{divisor, trueMonkey, falseMonkey}}
}
