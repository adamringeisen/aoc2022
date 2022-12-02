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

	score_map := map[string]int{"rock": 1, "paper": 2, "scissors": 3}
	game_map := map[string]string{"A": "rock", "B": "paper", "C": "scissors", "X": "lose", "Y": "draw", "Z": "win"}
	win := 6
	draw := 3

	if game_map[opp] == "rock" {
		if game_map[me] == "win" {
			return win + score_map["paper"]
		} else if game_map[me] == "draw" {
			return draw + score_map["rock"]
		} else {
			return score_map["scissors"]
		}
	} else if game_map[opp] == "paper" {
		if game_map[me] == "win" {
			return win + score_map["scissors"]
		} else if game_map[me] == "draw" {
			return draw + score_map["paper"]
		} else {
			return score_map["rock"]
		}
	} else if game_map[opp] == "scissors" {
		if game_map[me] == "win" {
			return win + score_map["rock"]
		} else if game_map[me] == "draw" {
			return draw + score_map["scissors"]
		} else {
			return score_map["paper"]
		}
	} else {
		panic("I don't know what's happening!!")
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
