package main

import (
	"fmt"
	"os"
	"bufio"
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

func main() {
	fmt.Println("Hello World")
}
