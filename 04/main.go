package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

type Grid struct {
	data []string
	h, w int
}

func NewGrid(lines []string) *Grid {
	return &Grid{
		data: lines,
		h:    len(lines),
		w:    len(lines[0]),
	}
}

func (g *Grid) InBounds(coord Coord) bool {
	return coord.r >= 0 && coord.r < g.h && coord.c >= 0 && coord.c < g.w
}

func (g *Grid) Get(coord Coord) byte {
	return g.data[coord.r][coord.c]
}

type Coord struct {
	r int
	c int
}

type Direction Coord

var (
	Up        = Direction{r: -1, c: 0}
	Down      = Direction{r: 1, c: 0}
	Left      = Direction{r: 0, c: -1}
	Right     = Direction{r: 0, c: 1}
	UpLeft    = Direction{r: -1, c: -1}
	UpRight   = Direction{r: -1, c: 1}
	DownLeft  = Direction{r: 1, c: -1}
	DownRight = Direction{r: 1, c: 1}
)

func (c Coord) Move(direction Direction) Coord {
	return Coord{c.r + direction.r, c.c + direction.c}
}

var nextChar = map[byte]byte{
	'X': 'M',
	'M': 'A',
	'A': 'S',
}

func solve1(grid *Grid, coord Coord, direction Direction, expectedChar byte) bool {
	if !grid.InBounds(coord) {
		return false
	}

	if grid.Get(coord) != expectedChar {
		return false
	}

	if grid.Get(coord) == 'S' {
		return true
	}

	next := nextChar[expectedChar]
	return solve1(grid, coord.Move(direction), direction, next)
}

func part1() {
	lines := ReadLines("input.txt")
	grid := NewGrid(lines)
	ans := 0
	for i := 0; i < grid.h; i++ {
		for j := 0; j < grid.w; j++ {
			start := Coord{i, j}
			for _, direction := range []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight} {
				if solve1(grid, start, direction, 'X') {
					ans++
				}
			}
		}
	}
	fmt.Println("Answer is", ans)
}

func part2() {
	lines := ReadLines("input.txt")
	grid := NewGrid(lines)
	ans := 0
	for i := 1; i < grid.h-1; i++ {
		for j := 1; j < grid.w-1; j++ {
			mid := Coord{i, j}
			masCount := 0
			if grid.Get(mid) != 'A' {
				continue
			}
			if grid.Get(mid.Move(UpLeft)) == 'M' && grid.Get(mid.Move(DownRight)) == 'S' {
				masCount++
			}
			if grid.Get(mid.Move(UpRight)) == 'M' && grid.Get(mid.Move(DownLeft)) == 'S' {
				masCount++
			}
			if grid.Get(mid.Move(DownRight)) == 'M' && grid.Get(mid.Move(UpLeft)) == 'S' {
				masCount++
			}
			if grid.Get(mid.Move(DownLeft)) == 'M' && grid.Get(mid.Move(UpRight)) == 'S' {
				masCount++
			}
			if masCount == 2 {
				ans++
			}
		}
	}
	fmt.Println("Answer is", ans)
}

func main() {
	var start time.Time
	start = time.Now()
	part1()
	fmt.Println("Part 1 finished in:", time.Since(start))

	start = time.Now()
	part2()
	fmt.Println("Part 2 finished in:", time.Since(start))
}
