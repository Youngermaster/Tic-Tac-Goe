package main

import (
	"fmt"
)

const (
	Empty = iota
	X
	O
)

type Board [3][3]int

func (b Board) String() string {
	var s string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch b[i][j] {
			case X:
				s += "X"
			case O:
				s += "O"
			default:
				s += " "
			}
			if j < 2 {
				s += " | "
			}
		}
		if i < 2 {
			s += "\n---------\n"
		}
	}
	return s
}

func (b Board) Play(player int, move int) error {
	if move < 1 || move > 9 {
		return fmt.Errorf("invalid move: %d", move)
	}
	i, j := (move-1)/3, (move-1)%3
	if b[i][j] != Empty {
		return fmt.Errorf("position already taken: %d", move)
	}
	b[i][j] = player
	return nil
}

func (b Board) CheckWin() int {
	for i := 0; i < 3; i++ {
		if b[i][0] != Empty && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return b[i][0]
		}
		if b[0][i] != Empty && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return b[0][i]
		}
	}
	if b[0][0] != Empty && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}
	if b[0][2] != Empty && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return b[0][2]
	}
	return Empty
}

func main() {
	var board Board
	var player int = X
	var winner int = Empty

	fmt.Println("Welcome to Tic Tac Go!")

	for {
		fmt.Println(board)
		fmt.Printf("Player %d's turn. Enter a number between 1 and 9: ", player)
		var move int
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// consume the leftover newline character
		fmt.Scanln()
		err = board.Play(player, move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		winner = board.CheckWin()
		if winner != Empty {
			break
		}
		if player == X {
			player = O
		} else {
			player = X
		}
	}

	fmt.Println(board)
	switch winner {
	case X:
		fmt.Println("Player X wins!")
	case O:
		fmt.Println("Player O wins!")
	default:
		fmt.Println("It's a tie!")
	}
}
