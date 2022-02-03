package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row, column int
}

type Board struct {
	rows      [5]int
	columns   [5]int
	positions map[int]Position
}

func newBoard() *Board {
	board := new(Board)
	board.positions = make(map[int]Position)

	return board
}

func (board *Board) add(n, row, column int) {
	board.rows[row] += n
	board.columns[column] += n
	board.positions[n] = Position{row, column}
}

func (board *Board) remove(n int) bool {
	if position, exists := board.positions[n]; exists {
		board.rows[position.row] -= n
		board.columns[position.column] -= n

    if board.rows[position.row]*board.columns[position.column] == 0 {
      return true
    }
	}
	return false
}

func (board *Board) sum() int {
	sum := 0
	for idx := range board.rows {
		sum += board.rows[idx]
	}
	return sum
}

type item struct {
  n int
  board *Board
}

var stack []item

func main() {
	file, err := os.Open("4/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	instructions := make([]int, 128)
	boards := make([]*Board, 128)
	board := newBoard()
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)

		switch len(args) {
		case 0:
			if row > 0 {
				boards = append(boards, board)
			}
			board, row = newBoard(), 0
		case 5:
			for column, str := range args {
				n, _ := strconv.Atoi(str)
				board.add(n, row, column)
			}
			row++
		default:
			args := strings.Split(args[0], ",")
			for _, value := range args {
				n, _ := strconv.Atoi(value)
				instructions = append(instructions, n)
			}
		}
	}
	if row > 0 {
		boards = append(boards, board)
	}

	for _, n := range instructions {
    i := 0
    for _, board := range boards {
      if board != nil && board.remove(n) {
        push(n, board)
        continue 
      }
      boards[i] = board
      i++
    }
    boards = boards[:i]
	}

  // Part one
  f, first := shift(stack)
  fmt.Println(f * first.sum())

  // Part two
  l, last := pop(stack)
  fmt.Println(l * last.sum())
}

func push(n int, board *Board ) {
  stack = append(stack, item{n, board})  
}

func shift(stack []item) (int, *Board) {
  first := stack[0]

  return first.n, first.board
}

func pop(stack []item) (int, *Board) {
  last := stack[len(stack) - 1]

  return last.n, last.board
}
