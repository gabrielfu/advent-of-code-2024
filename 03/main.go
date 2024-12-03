package main

import (
	"fmt"
	"os"
	"regexp"
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
	pattern := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	sum := 0
	for _, line := range lines {
		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			nums := strings.Split(match[4:len(match)-1], ",")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			sum += a * b
		}
	}
	fmt.Println("Answer is", sum)
}

func part2() {
	lines := ReadLines("input.txt")
	pattern := regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	sum := 0
	enabled := true
	for _, line := range lines {
		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				enabled = true
				continue
			} else if match == "don't()" {
				enabled = false
				continue
			}

			if !enabled {
				continue
			}
			nums := strings.Split(match[4:len(match)-1], ",")
			a, _ := strconv.Atoi(nums[0])
			b, _ := strconv.Atoi(nums[1])
			sum += a * b
		}
	}
	fmt.Println("Answer is", sum)
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
