package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Okay, I sort of cheated and didn't bother
// to parse the first bit of input into this map

var stacks = map[int]string{1: "WBDNCFJ", 2: "PZVQLST", 3: "PZBGJT", 4: "DTLJZBHC", 5: "GVBJS", 6: "PSQ", 7: "BVDFLMPN", 8: "PSMFBDLR", 9: "VDTR"}



func get_input() string {
	if len(os.Args) < 2 {
		panic("You must supply an input file.")
	} else {
		return os.Args[1]
	}
}

func get_lines() [][]string {
	input := get_input()
	data, _ := os.Open(input)
	var lines [][]string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		splitty := strings.Split(scanner.Text(), " ")
		lines = append(lines, []string{splitty[1],splitty[3], splitty[5]})

	}
	return lines
}


func get_last(s string) (string) {
	return s[len(s)-1:]
}

func trunc(s string) (string) {
	return s[:len(s)-1]
}

func move(times, s1, s2 int) {
	for i := 0; i < times; i++ {
		// get last of s1
		to_move := get_last(stacks[s1])
		// remove it from s1
		stacks[s1] = trunc(stacks[s1])
		// append it to s2
		stacks[s2] += to_move
	}
}

func main() {
	lines := get_lines()
	for _, line := range lines {
		times, _ := strconv.Atoi(line[0])
		s1, _ := strconv.Atoi(line[1])
		s2, _ := strconv.Atoi(line[2])
		move(times, s1, s2)
	}
	fmt.Println(stacks)
}
