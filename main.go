package main

import (
	"fmt"
)

type board struct {
	grid [][]rune
}

func main() {
	theBoard := newBoard()

	theBoard.printBoard()

}

func newBoard() board {
	newBoard := board{[][]rune{
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'}}}
	return newBoard
}

func (self *board) getPosition(i, j int) rune {
	return self.grid[i][j]
}

func (self *board) printBoard() {
	fmt.Printf("hello?! \n")
	fmt.Printf("   01234 \n")
	fmt.Printf("         \n")
	for j := range self.grid {
		fmt.Printf("%d ", 10+j)
		for i := range self.grid[0] {
			fmt.Printf("%c", self.getPosition(i, j))
		}
		fmt.Printf(" %d\n", 15+j)
	}
	fmt.Printf("         \n")
	fmt.Printf("   56789 \n")
}

// func (self *board) makeMove(pos int) {
// 	pos = 0
// 	putChar := self.grid[0][0]
// }
