package main

import (
	"fmt"
)

func printBoard() {
	i := 0
	fmt.Println("|=============================|")
	fmt.Println("|          TIC TAC GOE        |")
	fmt.Println("|=============================|")
	fmt.Printf("|    %d    |    %d    |    %d    |\n", i, i, i)
	fmt.Println("|-----------------------------|")
	fmt.Printf("|    %d    |    %d    |    %d    |\n", i, i, i)
	fmt.Println("|-----------------------------|")
	fmt.Printf("|    %d    |    %d    |    %d    |\n", i, i, i)
	fmt.Println("|=============================|")
}

func main() {
	boardValues := [3][3]uint8{}
	fmt.Println(boardValues)
	fmt.Println(len(boardValues))
	fmt.Println("Let's Get started!")
	printBoard()
}