package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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

func part1() {
	lines := ReadLines("input.txt")

	var left sort.IntSlice
	var right sort.IntSlice
	for _, line := range lines {
		split := strings.Fields(line)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		left = append(left, l)
		right = append(right, r)
	}
	left.Sort()
	right.Sort()

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(right[i] - left[i])))
	}
	fmt.Println("Answer is", sum)
}

func part2() {
	lines := ReadLines("input.txt")

	var left sort.IntSlice
	var right sort.IntSlice
	for _, line := range lines {
		split := strings.Fields(line)
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])
		left = append(left, l)
		right = append(right, r)
	}

	rightCounter := make(map[int]int)
	for _, r := range right {
		rightCounter[r]++
	}

	similarity := 0
	for _, l := range left {
		similarity += l * rightCounter[l]
	}
	fmt.Println("Answer is", similarity)
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
