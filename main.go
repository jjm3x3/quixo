package main

import (
	"fmt"
	_ "strconv"
)

type board struct {
	grid [][]rune
	turn bool
}

func main() {
	theBoard := newBoard()

	theBoard.printBoard()

	theBoard.makeMove(0)
	theBoard.makeMove(0)
	theBoard.makeMove(0)
	theBoard.makeMove(0)
	// theBoard.makeMove(0)
	// theBoard.makeMove(1)
	// theBoard.makeMove(1)
	// theBoard.makeMove(1)
	// theBoard.makeMove(1)
	// theBoard.makeMove(1)
	// theBoard.makeMove(5)
	// theBoard.makeMove(10)
	// theBoard.makeMove(10)
	// theBoard.makeMove(11)
	// theBoard.makeMove(15)

	// some one should win
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	theBoard.makeMove(0)

	theBoard.makeMove(1)

	theBoard.printBoard()

}

func newBoard() board {
	newBoard := board{grid: [][]rune{
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'}}, turn: true}
	return newBoard
}

// func checkForWin() bool {
// 	for
// }

func (self *board) getPosition(i, j int) rune {
	return self.grid[i][j]
}

func (self *board) printBoard() {
	fmt.Printf("         \n")
	fmt.Printf("   01234 \n")
	fmt.Printf("         \n")
	for j := range self.grid {
		fmt.Printf("%d ", 10+j)
		for i := range self.grid[0] {
			fmt.Printf("%c", self.getPosition(j, i))
		}
		fmt.Printf(" %d\n", 15+j)
	}
	fmt.Printf("         \n")
	fmt.Printf("   56789 \n")
	fmt.Printf("         \n")
}

func (self *board) makeMove(pos int) {
	var theMove rune
	if self.turn {
		theMove = 'X'
	} else {
		theMove = 'O'
	}
	self.turn = !self.turn

	if pos >= 0 && pos < 5 {

		for row := 0; row < 5; row++ {
			restOfRow := self.grid[row][pos+1 : len(self.grid[0])]
			startOfRow := self.grid[row][0:pos]
			var newRow []rune
			newRow = append(newRow, startOfRow...)
			if row == 4 {
				newRow = append(newRow, theMove)
			} else {
				newPeice := self.grid[row+1][pos]
				newRow = append(newRow, newPeice)
			}
			newRow = append(newRow, restOfRow...)
			self.grid[row] = newRow
		}

	}
	if pos >= 5 && pos < 10 {
		pos = pos - 5

		for row := 4; row >= 0; row-- {
			restOfRow := self.grid[row][pos+1 : len(self.grid[0])]
			startOfRow := self.grid[row][0:pos]
			var newRow []rune
			newRow = append(newRow, startOfRow...)
			if row == 0 {
				newRow = append(newRow, theMove)
			} else {
				newPeice := self.grid[row-1][pos]
				newRow = append(newRow, newPeice)
			}
			newRow = append(newRow, restOfRow...)
			self.grid[row] = newRow
		}
	}

	if pos >= 10 && pos < 15 {
		pos = pos - 10

		// fmt.Printf("%v\n", self.grid[thing])
		var newRow []rune
		newRow = append(newRow, self.grid[pos][1:len(self.grid)]...)
		newRow = append(newRow, theMove) //self.grid[pos][0])
		self.grid[pos] = newRow

	}

	if pos >= 15 && pos < 20 {
		pos = pos - 15

		var newRow []rune
		newRow = append(newRow, theMove) //self.grid[pos][0])
		newRow = append(newRow, self.grid[pos][0:len(self.grid)-1]...)
		self.grid[pos] = newRow

	}

}
