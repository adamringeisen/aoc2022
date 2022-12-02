package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func score(opp string, me string) int {

	score_map := map[string]int{"X": 1, "Y": 2, "Z": 3}
	game_map := map[string]string{"A": "rock", "B": "paper", "C": "scissors", "X": "rock", "Y": "paper", "Z": "scissors"}
	win := score_map[me] + 6
	draw := score_map[me] + 3
	loss := score_map[me]
	if game_map[opp] == game_map[me] {
		return draw
	} else {
		if game_map[opp] == "rock" {
			if game_map[me] == "paper" {
				return win
			} else {
				return loss
			}
		} else if game_map[opp] == "paper" {
			if game_map[me] == "scissors" {
				return win
			} else {
				return loss
			}
		} else if game_map[opp] == "scissors" {
			if game_map[me] == "rock" {
				return win
			} else {
				return loss
			}
		} else {
			panic("I don't know what's happening!!")
		}
	}
}

func main() {
	final_score := 0
	lines := get_lines()
	for _, line := range lines {
		final_score += score(string(line[0]), string(line[2]))
	}
	fmt.Printf("The final score is %d\n", final_score)
}
