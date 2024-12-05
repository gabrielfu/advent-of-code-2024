package main

import (
	"fmt"
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

func IndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 // not found.
}

type Rule struct {
	first  int
	second int
}

type Update struct {
	pages []int
}

func (u Update) Copy() Update {
	pages := make([]int, len(u.pages))
	copy(pages, u.pages)
	return Update{pages}
}

func (u *Update) SatisfyRule(r Rule) bool {
	first := IndexOf(r.first, u.pages)
	second := IndexOf(r.second, u.pages)
	return first == -1 || second == -1 || first < second
}

func (u *Update) SatisfyAllRules(rules []Rule) bool {
	for _, r := range rules {
		if !u.SatisfyRule(r) {
			return false
		}
	}
	return true
}

func (u *Update) GetMiddle() int {
	length := len(u.pages) // always odd
	return u.pages[length/2]
}

func (u Update) Fix(rules []Rule) Update {
	new := u.Copy()
	for _, r := range rules {
		first := IndexOf(r.first, new.pages)
		second := IndexOf(r.second, new.pages)
		if first != -1 && second != -1 && first > second {
			new.pages[first], new.pages[second] = new.pages[second], new.pages[first]
		}
	}
	return new
}

type Puzzle struct {
	updates []Update
	rules   []Rule
}

func ParseInput(lines []string) Puzzle {
	p := Puzzle{}
	firstSection := true
	for _, line := range lines {
		if line == "" {
			firstSection = false
			continue
		}

		if firstSection {
			split := strings.Split(line, "|")
			first, _ := strconv.Atoi(split[0])
			second, _ := strconv.Atoi(split[1])
			p.rules = append(p.rules, Rule{first, second})
		} else {
			split := strings.Split(line, ",")
			pages := make([]int, len(split))
			for i, s := range split {
				pages[i], _ = strconv.Atoi(s)
			}
			p.updates = append(p.updates, Update{pages})
		}
	}
	return p
}

func part1() {
	lines := ReadLines("input.txt")
	p := ParseInput(lines)
	ans := 0
	for _, u := range p.updates {
		if u.SatisfyAllRules(p.rules) {
			ans += u.GetMiddle()
		}
	}
	fmt.Println("Answer is", ans)
}

func part2() {
	lines := ReadLines("input.txt")
	p := ParseInput(lines)
	ans := 0
	for _, u := range p.updates {
		if !u.SatisfyAllRules(p.rules) {
			for {
				u = u.Fix(p.rules)
				if u.SatisfyAllRules(p.rules) {
					break
				}
			}
			ans += u.GetMiddle()
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
