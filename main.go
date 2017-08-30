package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	board := setup(4)
	for has(board, 0) {
		addNumber(board)
		display(board)
		switch getOption() {
		case 'd':
			shiftLeft(board)
		case 's':
			shiftRight(board)
		case 'a':
			shiftUp(board)
		case 'w':
			shiftDown(board)
		}
		if has(board, 2048) {
			fmt.Println("Congradulations you got 2048")
			break
		}
	}
	fmt.Println("done")
}

// addNumber assumes there is an empty location
func addNumber(board [][]int) {
	options := []int{2, 2, 2, 4}
	x, y := 0, 0
	for {
		x, y = rand.Int()%len(board), rand.Int()%len(board)
		if board[x][y] == 0 {
			break
		}
	}
	board[x][y] = options[rand.Int()%len(options)]
}

func has(board [][]int, val int) bool {
	for _, row := range board {
		for _, tile := range row { // column
			if tile == val {
				return true
			}
		}
	}
	return false
}

func getOption() rune {
	var char rune
	for {
		if fmt.Scanf("%c", &char); strings.ContainsRune("dsaw", char) {
			break
		}
	}
	return char
}

func setup(dim int) [][]int {
	board := [][]int{}
	for i := 0; i < dim; i++ {
		row := []int{}
		for j := 0; j < dim; j++ {
			row = append(row, 0)
		}
		board = append(board, row)
	}
	return board
}

func display(board [][]int) {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	for _, row := range board {
		for _, val := range row {
			if val != 0 {
				fmt.Printf("[%4d]", val)
			} else {
				fmt.Print("[    ]")
			}
		}
		fmt.Println()
	}
	fmt.Print(": ")
}

func slideTile(board [][]int, toI, toJ, fromI, fromJ int) {
	if board[toI][toJ] == board[fromI][fromJ] {
		board[toI][toJ] += board[fromI][fromJ]
		board[fromI][fromJ] = 0
	} else if board[toI][toJ] == 0 {
		board[toI][toJ] = board[fromI][fromJ]
		board[fromI][fromJ] = 0
	}
}

func shiftRight(board [][]int) {
	for i := range board {
		for step := len(board[i]) - 1; step > 0; step-- {
			for j := len(board[i]) - 1; j > 0; j-- {
				slideTile(board, i, j, i, j-1)
			}
		}
	}
}

func shiftLeft(board [][]int) {
	for i := range board {
		for step := len(board[i]) - 1; step > 0; step-- {
			for j := 0; j < len(board[i])-1; j++ {
				slideTile(board, i, j, i, j+1)
			}
		}
	}
}

func shiftUp(board [][]int) {
	i := 0
	for j := range board[i] {
		for step := len(board) - 1; step > 0; step-- {
			for i := 0; i < len(board)-1; i++ {
				slideTile(board, i, j, i+1, j)
			}
		}
	}
}

func shiftDown(board [][]int) {
	i := 0
	for j := range board[i] {
		for step := len(board) - 1; step > 0; step-- {
			for i := len(board) - 1; i >= 1; i-- {
				slideTile(board, i, j, i-1, j)
			}
		}
	}
}
