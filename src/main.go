package main

import (
	"fmt"
)

func printBoard() {
	i := 10
	fmt.Println("|=============================|")
	fmt.Println("|         TIC TAC GOE         |")
	fmt.Println("|=============================|")
	fmt.Sprintf("|    %d     |         |         |", i)
	fmt.Println("|-----------------------------|")
	fmt.Println("|         |         |         |")
	fmt.Println("|-----------------------------|")
	fmt.Println("|         |         |         |")
	fmt.Println("|=============================|")
}

func main() {
	boardValues := [3][3]uint8{}
	fmt.Println(boardValues)
	fmt.Println(len(boardValues))
	fmt.Println("Let's Get started!")
	printBoard()
}