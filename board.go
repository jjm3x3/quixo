package main

import (
	"errors"
	"fmt"
)

type board struct {
	grid [][]rune
	turn bool
}

func (self *board) whoseTurn() rune {
	if self.turn {
		return 'X'
	} else {
		return 'O'
	}
}

func newBoard(layout [][]rune) *board {
	var newBoard board
	if layout != nil {
		newBoard = board{grid: layout, turn: true}
	} else {
		newBoard = board{grid: [][]rune{
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'},
			[]rune{'#', '#', '#', '#', '#'}}, turn: true}
	}
	return &newBoard
}

func (self *board) checkForWin() rune {
	leads := make([]rune, 12)

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

	result := '#'

	for i := range leads {
		// fmt.Printf("What does this look like: %c\n", leads[i])
		if leads[i] != '#' {
			if self.whoseTurn() == leads[i] {
				return leads[i]
			} else {
				result = leads[i]
			}
		}

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

func isBottom(move int) bool {
	return move >= 5 && move < 10
}

func isTop(move int) bool {
	return move >= 0 && move < 5
}

func getTargetPeice(pos, dest int) (int, int) {
	var row int
	var col int
	if isTop(pos) || isBottom(pos) {
		col = pos
		row = 0
		if isBottom(pos) {
			col = pos - 5
			row = 4
		}
	} else { // check for pos and dest same TODO
		col = 0
		row = pos - 10
		if isRight(pos) {
			row = pos - 15
			col = 4
		}
	}
	return row, col
}

func (self *board) makeMove(pos, dest int) error {
	// fmt.Printf("move to?: %v\n", pos)
	// fmt.Printf("wats the move %c\n", self.whoseTurn())

	sanityCheck(pos, dest) // not an ideal location but this is
	//about as good as it gets right now
	var row int
	if isTop(pos) || isBottom(pos) {
		row, col := getTargetPeice(pos, dest)
		// fmt.Printf("where are we moving? (%v,%v)\n", row, col)

		if self.checkValidPeiceSelection(row, col) {
			return errors.New("that is an illegal Move")
		}

		if isBottom(dest) || isTop(dest) {
			self.cycleColumn(pos, row, col, dest)

			// return nil // don't know why this can't return nil
		} else { // this is for starts on the top and bottom
			// and a move right or left
			self.cycleRow(row, col, dest)
			return nil
		}

	} else { // check for pos and dest same TODO
		row, col := getTargetPeice(pos, dest)
		// fmt.Printf("which row? %v\n", row)
		if self.checkValidPeiceSelection(row, col) {
			return errors.New("that is an illegal Move")
		}
		if isBottom(dest) || isTop(dest) {
			self.cycleColumn(pos, row, col, dest)
		} else {
			self.cycleRow(row, col, dest)
		}

	}

	// switch the turn
	self.turn = !self.turn
	return nil
}

func (self *board) cycleRow(row, col, dest int) {
	var newRow []rune
	if isLeft(dest) {
		newRow = append(newRow, self.whoseTurn())
	}
	newRow = append(newRow, self.grid[row][0:col]...)
	newRow = append(newRow, self.grid[row][col+1:len(self.grid[row])]...)
	if isRight(dest) {
		newRow = append(newRow, self.whoseTurn())
	}
	self.grid[row] = newRow
}

func (self *board) cycleColumn(pos, moveRow, col, dest int) {
	var row int
	if moveRow != 4 {
		row = moveRow
	}
	for ; row < 5; row++ {
		rowCur := row
		if isTop(dest) {
			rowCur = 4 - row
			// fmt.Printf("oporate on which row: %v\n", row)
		}
		restOfRow := self.grid[rowCur][col+1 : len(self.grid[0])]
		startOfRow := self.grid[rowCur][0:col]
		var newRow []rune
		newRow = append(newRow, startOfRow...)
		if (rowCur == 4 && isBottom(dest)) || (rowCur == 0 && isTop(dest)) {
			newRow = append(newRow, self.whoseTurn())
		} else {
			var newPeice rune
			if isTop(dest) {
				newPeice = self.grid[rowCur-1][col]
			} else {
				newPeice = self.grid[rowCur+1][col]
			}
			newRow = append(newRow, newPeice)
		}
		newRow = append(newRow, restOfRow...)
		self.grid[rowCur] = newRow
	}
}

func (self *board) checkValidPeiceSelection(row, col int) bool {
	// fmt.Printf("Peice from (%v,%v)\n", row, col)
	selectedPeice := self.grid[row][col]
	// fmt.Printf("what is this thing: %c\n", selectedPeice)
	return selectedPeice != self.whoseTurn() && selectedPeice != '#'
}

func isLeft(move int) bool {
	return move >= 10 && move < 15
}

func isRight(move int) bool {
	return move >= 15 && move < 20
}

func (self *board) equals(other *board) bool {
	for i := 0; i < len(self.grid); i++ {
		for j := 0; j < len(self.grid[0]); j++ {
			if self.getPosition(i, j) != other.getPosition(i, j) {
				return false
			}
		}
	}
	return true
}

func sanityCheck(pos, dest int) error {
	if isRight(dest) {
		if pos == 4 || pos == 9 {
			return errors.New("This move is nonsesical")
		}

	}
	if isLeft(dest) {
		if pos == 0 || pos == 5 {
			return errors.New("This move is nonsesical")
		}
	}
	return nil
}
