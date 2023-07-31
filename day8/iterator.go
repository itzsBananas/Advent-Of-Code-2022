package main

type coord struct {
	x, y         int
	val          int
	newIteration bool
}

func TopLeft(grid [][]int) <-chan coord {
	c := make(chan coord)
	go func() {
		for i := 0; i < len(grid); i++ {
			newI := true
			for j := 0; j < len(grid[i]); j++ {
				c <- coord{i, j, grid[i][j], newI}
				newI = false
			}
		}
		close(c)
	}()
	return c
}

func TopRight(grid [][]int) <-chan coord {
	c := make(chan coord)
	go func() {
		for j := 0; j < len(grid[0]); j++ {
			newI := true
			for i := len(grid) - 1; i >= 0; i-- {
				c <- coord{i, j, grid[i][j], newI}
				newI = false
			}
		}
		close(c)
	}()
	return c
}

func BottomLeft(grid [][]int) <-chan coord {
	c := make(chan coord)
	go func() {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			newI := true
			for i := 0; i < len(grid); i++ {
				c <- coord{i, j, grid[i][j], newI}
				newI = false
			}
		}
		close(c)
	}()
	return c
}

func BottomRight(grid [][]int) <-chan coord {
	c := make(chan coord)
	go func() {
		for i := len(grid) - 1; i >= 0; i-- {
			newI := true
			for j := len(grid[i]) - 1; j >= 0; j-- {
				c <- coord{i, j, grid[i][j], newI}
				newI = false
			}
		}
		close(c)
	}()
	return c
}
