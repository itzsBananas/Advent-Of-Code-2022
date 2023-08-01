package main

import (
	"bytes"
	"fmt"
)

type grid map[string]bool

type pos struct {
	x, y int
}

type bridge struct {
	grid
	knots []pos
}

func NewBridge(s int) *bridge {
	var b bridge
	b.grid = make(grid)
	b.knots = make([]pos, s)
	b.SavePosition(pos{0, 0})
	return &b
}

func (b *bridge) MoveVertical(s int) {
	if s > 0 {
		b.knots[0].y += 1
		b.resolveGap(1)
		b.MoveVertical(s - 1)
	} else if s < 0 {
		b.knots[0].y -= 1
		b.resolveGap(1)
		b.MoveVertical(s + 1)
	}
}

func (b *bridge) MoveHorizontal(s int) {
	if s > 0 {
		b.knots[0].x += 1
		b.resolveGap(1)
		b.MoveHorizontal(s - 1)
	} else if s < 0 {
		b.knots[0].x -= 1
		b.resolveGap(1)
		b.MoveHorizontal(s + 1)
	}
}

func (b *bridge) SavePosition(p pos) {
	b.grid[fmt.Sprintf("(%d,%d)", p.x, p.y)] = true
}

func (b *bridge) HasPosition(p pos) bool {
	return b.grid[fmt.Sprintf("(%d,%d)", p.x, p.y)]
}

// not dynamic to size, but okay for testing small boards
func (b bridge) String() string {
	var buf bytes.Buffer
	for y := 5; y >= -5; y-- {
	loop:
		for x := -5; x <= 5; x++ {
			for i, k := range b.knots {
				if k.x == x && k.y == y {
					buf.WriteString(fmt.Sprintf("%d", i+1))
					continue loop
				}
			}
			if x == 0 && y == 0 {
				buf.WriteRune('S')
				continue loop
			}
			if b.HasPosition(pos{x, y}) {
				buf.WriteRune('#')
			} else {
				buf.WriteRune('.')
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
	// return fmt.Sprintf("%v", b.grid)
}

func (b bridge) Count() int {
	return len(b.grid)
}

// assume p2 needs to be moved closer to p1 if necessary
func (b *bridge) resolveGap(ind int) {
	if ind >= len(b.knots) {
		return
	}
	p1 := &b.knots[ind-1]
	p2 := &b.knots[ind]
	xDiff := p1.x - p2.x
	if xDiff < 0 {
		xDiff = -xDiff
	}
	yDiff := p1.y - p2.y
	if yDiff < 0 {
		yDiff = -yDiff
	}
	if xDiff < 2 && yDiff < 2 {
		return
	}

	if xDiff < 2 {
		p2.x = p1.x
	} else {
		p2.x = (p1.x + p2.x) / 2
	}
	if yDiff < 2 {
		p2.y = p1.y
	} else {
		p2.y = (p1.y + p2.y) / 2
	}
	if ind == len(b.knots)-1 {
		b.SavePosition(*p2)
	}
	b.resolveGap(ind + 1)
}
