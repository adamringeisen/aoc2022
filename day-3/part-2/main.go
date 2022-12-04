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




func group_input(s []string) ([][]string) {
	var grouped [][]string
	for i := 0; i < len(s); i += 3 {
		grouped = append(grouped, []string{s[i],s[i+1],s[i+2]})
	}
	return grouped
}


func check_shared(s []string) (rune) {
	for _, char := range s[0] {
		if strings.Contains(s[1], string(char)) && strings.Contains(s[2], string(char))  {
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
	result := 0
	input := get_lines()
	fmt.Printf("Input length: %d\nDivided by 3: %d\n", len(input), len(input)/3)
	lines := group_input(input)
	fmt.Printf("Grouped Input length: %d\n", len(lines))
	for _, char := range lines {
		result += get_value((check_shared(char)))
		
	}
	fmt.Printf("The result is: %d\n", result)
}
