package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func get_input() string {
	if len(os.Args) < 2 {
		panic("You must supply an input file.")
	} else {
		return os.Args[1]
	}
}

func get_lines() []string {
	input := get_input()
	data, err := os.Open(input)
	check(err)
	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}



func split_string(s string) ([]string) {
	return []string{s[:len(s)/2], s[len(s)/2:]}
}

func check_shared(s []string) (rune) {
	for _, char := range s[0] {
		if strings.Contains(s[1], string(char)) {
			return rune(char)
		}
	}
	return 33		// return "!" if reached
}

func get_value(r rune) (int) {
	if r < 91 {
		return int(r) - 38
	} else {
		return int(r) - 96
	}
}
func main() {
	final_score := 0
	lines := get_lines()
	for _, line := range lines {
		split := split_string(line)
		final_score += get_value(check_shared(split))
	}
	fmt.Printf("The final score is %d\n", final_score)
}
