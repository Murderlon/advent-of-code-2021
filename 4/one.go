package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	contents, err := ioutil.ReadFile("4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(contents), "\n\n")
	instructions := strings.Split(input[0], ",")
	boards := matrixify(input[1:])

	walk(instructions, boards)
}

func walk(instructions []string, boards [][][]string) {
	for _, instruction := range instructions {
		for boardIdx, board := range boards {
			for rowIdx, row := range board {
				for charIdx, char := range row {
					if char == instruction {
						boards[boardIdx][rowIdx][charIdx] = "x"

						// rows
						if every(row, func(str string) bool { return str == "x" }) {
							i, _ := strconv.Atoi(instruction)
							u := unmarked(board)
							fmt.Println(i, "*", u, "=", i*u)
							return
						}

						// columns
						col := 0
						marked := 0

						for col < len(board[0]) {
							for _, r := range board {
								if len(r) == 0 {
									continue
								}
								if r[col] == "x" {
									marked++
									if marked == len(boards[0]) {
										i, _ := strconv.Atoi(instruction)
										u := unmarked(board)
										fmt.Println(i, "*", u, "=", i*u)
										return
									}
								}
							}
							marked = 0
							col++
						}
					}
				}
			}
		}
	}
}

func unmarked(board [][]string) int {
	total_unmarked := 0

	for _, r := range board {
		for _, c := range r {
			int, err := strconv.Atoi(c)
			if err != nil {
				continue
			}
			total_unmarked += int
		}
	}

	return total_unmarked
}

func every(arr []string, predicate func(string) bool) bool {
	bool := true

	for _, value := range arr {
		if !predicate(value) {
			bool = false
		}
	}

	return bool
}

func matrixify(boards []string) [][][]string {
	matrix := make([][][]string, len(boards))

	for index, board := range boards {
		rows := strings.Split(board, "\n")
		for rowIdx, row := range rows {
			if matrix[index] == nil {
				matrix[index] = make([][]string, len(rows))
			}
			matrix[index][rowIdx] = strings.Fields(row)
		}
	}

	return matrix
}
