package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var per_elf = 0
var richest_elf = 0

func check(e error) {
	if e != nil {
		panic(e)
	}
}




func main() {
	
	dat, err := os.Open("day_1_input.txt")
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if per_elf > richest_elf {
				richest_elf = per_elf
			}
			per_elf = 0
		} else {
			calories, err := strconv.Atoi(line)
			check(err)
			per_elf += calories
		}
	}
	fmt.Printf("The Elf with the most calories has %d of them\n", richest_elf)
}
