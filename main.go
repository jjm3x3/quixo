package main

import (
	_ "bytes"
	"errors"
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

	// theBoard.makeMove(0)
	// theBoard.makeMove(0)
	// theBoard.makeMove(0)
	// theBoard.makeMove(0)

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
	aWin := theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 0) //x go
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 1) //o go
	// 2
	// theBoard.printBoard()
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 0) //x go
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 1)
	//4
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 0)
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 1)
	// 6
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 0)
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 1)
	// 8
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 0)
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)
	tryMove(theBoard, 1)
	// 10
	aWin = theBoard.checkForWin()
	fmt.Printf("did someone win?!: %v\n", aWin)

	// some one should win
	// tryMove(theBoard, 10) //x go
	// tryMove(theBoard, 11) //o go
	// // 2
	// // theBoard.printBoard()
	// tryMove(theBoard, 10) //x go
	// tryMove(theBoard, 11) //o go
	// //4
	// tryMove(theBoard, 10) //x go
	// tryMove(theBoard, 11) //o go
	// // 6
	// tryMove(theBoard, 10) //x go
	// tryMove(theBoard, 11) //o go
	// // 8
	// tryMove(theBoard, 10) //x go
	// tryMove(theBoard, 11) //o go

	//should be illegal
	// tryMove(theBoard, 0) //x go
	// tryMove(theBoard, 5) //o go

	// tryMove(theBoard, 5) //x go
	// tryMove(theBoard, 0) //o go

	// tryMove(theBoard, 10) //x go
	// tryMove(theBoard, 15) //o go

	// tryMove(theBoard, 15) //x go
	// tryMove(theBoard, 10) //o go

	theBoard.printBoard()

}

func tryMove(theBoard *board, x int) {
	// fmt.Printf("whose turn? %v\n", theBoard.turn)
	err := theBoard.makeMove(x)
	if err != nil {
		panic("ilegal mOVE!")
	}
}

func newBoard() *board {
	newBoard := board{grid: [][]rune{
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'}}, turn: true}
	return &newBoard
}

func (self *board) checkForWin() bool {
	leads := make([]rune, 12)
	// wins := make([]bool, 12)

	// for i := 0; i < len(self.board); i++ {
	// 	if i == 0 {
	// 		if slef.board[0][i] != '#' {
	// 			leads[11] = self.board[0][i]
	// 		} else {
	// 			wins[11] = true
	// 		}
	// 	}
	// 	if slef.board[0][i] != '#' {
	// 		leads[i] = self.board[0][i]
	// 	} else {
	// 		wins[i] = true
	// 	}
	// }

	// for i := 0; i < len(self.board); i++ {
	// 	if slef.board[i][0] != '#' {
	// 		leads[i+5] = self.board[0][i]
	// 	} else {
	// 		wins[i+5] = true
	// 	}
	// 	if i == 5 {
	// 		if slef.board[0][i] != '#' {
	// 			leads[12] = self.board[0][i]
	// 		} else {
	// 			wins[12] = true
	// 		}
	// 	}
	// }

	for i := 0; i < len(self.grid); i++ {
		for j := 0; j < len(self.grid); j++ {
			if i == 0 && j == 0 {
				leads[0] = self.grid[i][j]
				leads[5] = self.grid[i][j]
				leads[10] = self.grid[i][j]
			} else if i == 0 {
				if j == 4 {
					leads[11] = self.grid[i][j]
				}
				leads[j+5] = self.grid[i][j]
				if leads[0] != self.grid[i][j] {
					leads[0] = '#'
				}
			} else if j == 0 {
				leads[i] = self.grid[i][j]
				if leads[j+5] != self.grid[i][j] {
					leads[j+5] = '#'
				}
			} else {
				if leads[i] != self.grid[i][j] {
					leads[i] = '#'
				}
				if leads[j+5] != self.grid[i][j] {
					leads[j+5] = '#'
				}
				if i == j {
					if leads[10] != self.grid[i][j] {
						leads[10] = '#'
					}

				}
			}
			if i+j == 4 {
				if leads[11] != self.grid[i][j] {
					leads[11] = '#'
				}

			}
			// fmt.Printf("what is the val of %c\n", leads)
		}
	}

	result := false

	for i := range leads {
		isWinner := leads[i] != '#'
		// fmt.Printf("is it true now?: %v\n", isWinner)
		result = result || isWinner
	}

	return result
}

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

func (self *board) makeMove(pos int) error {
	var theMove rune
	if self.turn {
		theMove = 'X'
	} else {
		theMove = 'O'
	}

	if pos >= 0 && pos < 5 {

		for row := 0; row < 5; row++ {
			restOfRow := self.grid[row][pos+1 : len(self.grid[0])]
			oldPeice := self.grid[row][pos]
			if oldPeice != theMove && oldPeice != '#' {
				fmt.Printf("what is this thing: %c\n", oldPeice)
				return errors.New("that is an illegal Move")
			}
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
			oldPeice := self.grid[row][pos]
			if oldPeice != theMove && oldPeice != '#' {
				fmt.Printf("what is this thing: %c\n", oldPeice)
				return errors.New("that is an illegal Move")
			}
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
		oldPeice := self.grid[pos][0]
		if oldPeice != theMove && oldPeice != '#' {
			fmt.Printf("what is this thing: %c\n", oldPeice)
			return errors.New("that is an illegal Move")
		}
		newRow = append(newRow, self.grid[pos][1:len(self.grid)]...)
		newRow = append(newRow, theMove) //self.grid[pos][0])
		self.grid[pos] = newRow

	}

	if pos >= 15 && pos < 20 {
		pos = pos - 15

		var newRow []rune
		oldPeice := self.grid[pos][len(self.grid)]
		if oldPeice != theMove && oldPeice != '#' {
			fmt.Printf("what is this thing: %c\n", oldPeice)
			return errors.New("that is an illegal Move")
		}
		newRow = append(newRow, theMove) //self.grid[pos][0])
		newRow = append(newRow, self.grid[pos][0:len(self.grid)-1]...)
		self.grid[pos] = newRow

	}

	// switch the turn
	self.turn = !self.turn
	return nil
}
