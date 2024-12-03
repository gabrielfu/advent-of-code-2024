package main

import (
	"fmt"
	"math"
	"os"
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

func isSafe(nums []int) bool {
	increase, decrease := 0, 0
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		prev := nums[i-1]

		absDiff := int(math.Abs(float64(cur - prev)))
		if absDiff == 0 || absDiff > 3 {
			return false
		}

		if cur > prev {
			increase++
		} else {
			decrease++
		}

		if increase > 0 && decrease > 0 {
			return false
		}
	}
	return true
}

func part1() {
	lines := ReadLines("input.txt")

	numSafe := 0
	for _, line := range lines {
		var nums []int
		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		if isSafe(nums) {
			numSafe++
		}
	}
	fmt.Println("Answer is", numSafe)
}

func part2() {
	lines := ReadLines("input.txt")

	numSafe := 0
	for _, line := range lines {
		var nums []int
		for _, num := range strings.Fields(line) {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		for i := 0; i < len(nums); i++ {
			removed := append([]int(nil), nums[:i]...)
			removed = append(removed, nums[i+1:]...)
			if isSafe(removed) {
				numSafe++
				break
			}
		}
	}
	fmt.Println("Answer is", numSafe)
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
