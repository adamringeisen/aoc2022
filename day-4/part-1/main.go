package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() string {
	if len(os.Args) < 2 {
		panic("You must supply an input file.")
	} else {
		return os.Args[1]
	}
}

func get_lines() [][]string {
	input := get_input()
	data, err := os.Open(input)
	check(err)
	var lines [][]string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		splitty := strings.Split(scanner.Text(), ",")
		lines = append(lines, splitty)

	}
	return lines
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func just_answer(amin, amax, bmin, bmax int) bool {
	// find larger range
	arange := amax - amin
	brange := bmax - bmin
	// if a is bigger
	if arange > brange {
		if bmin >= amin && bmax <= amax {
			//			fmt.Printf("A > B [%d, %d], [%d, %d] True\n", amin, amax, bmin, bmax) 
			return true
		}
	} else
	// if b is bigger
	{
		if amin >= bmin && amax <= bmax {
			// fmt.Printf("B > A [%d, %d], [%d, %d] True\n", amin, amax, bmin, bmax)
			return true
		}
	}
	// fmt.Printf("AR: %d, BR: %d - [%d, %d], [%d, %d] False\n", arange, brange,amin, amax, bmin, bmax)
	return false
}

func cv(s string) int {
	a, e := strconv.Atoi(s)
	check(e)
	return a
}

func main() {
	lines := get_lines()
	result := 0

	for _, line := range lines {
		a := strings.Split(line[0], "-")
		b := strings.Split(line[1], "-")

		if just_answer(cv(a[0]), cv(a[1]), cv(b[0]), cv(b[1])) {
			result++
		}
	}
	fmt.Printf("Result is: %d\n", result)

}
