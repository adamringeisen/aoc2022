package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)
var per_elf = 0
var richest_elf = 0
var all_elves []int

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
			all_elves = append(all_elves, per_elf)
			per_elf = 0
		} else {
			calories, err := strconv.Atoi(line)
			check(err)
			per_elf += calories
		}
	}
	sort.Ints(all_elves)
	slc := all_elves[len(all_elves)-3:]
	
	fmt.Printf("The total of the top three elves is %d\n", slc[0] + slc[1] + slc[2])
}
