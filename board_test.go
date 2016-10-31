package main

import (
	_ "fmt"
	"testing"
)

func TestMoves(t *testing.T) {

	theBoard := newBoard(nil)

	theBoard.printBoard()

	theBoard.makeMove(0, 5)
	if theBoard.getPosition(4, 0) != 'X' {
		t.Errorf("this should be a plenty fine opening move")
	}
	err := theBoard.makeMove(0, 5)
	if err != nil {
		t.Logf("this is a bad move?!")
	}
	if theBoard.getPosition(4, 0) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(0, 5)
	if theBoard.getPosition(4, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(0, 5)
	if theBoard.getPosition(4, 0) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}

	theBoard.makeMove(0, 5)
	if theBoard.getPosition(4, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1, 6)
	if theBoard.getPosition(4, 1) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1, 6)
	if theBoard.getPosition(4, 1) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1, 6)
	if theBoard.getPosition(4, 1) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1, 6)
	if theBoard.getPosition(4, 1) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1, 6)
	if theBoard.getPosition(4, 1) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	err = theBoard.makeMove(5, 0)
	if theBoard.getPosition(0, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	if err != nil {
		t.Errorf("This should be plenty valid")
	}
	if theBoard.getPosition(4, 0) != 'O' {
		t.Errorf("some move messed up!")
	}
	err = theBoard.makeMove(10, 15)
	if err == nil {
		t.Errorf("this shouldent be allowed")
	}

	// make expicit turn change for testing
	theBoard.turn = !theBoard.turn

	err = theBoard.makeMove(10, 15)
	if err != nil {
		t.Errorf("This should be just fine")
	}
	if theBoard.getPosition(0, 4) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	err = theBoard.makeMove(12, 17)
	if err != nil {
		t.Errorf("This should be just fine")
	}
	if theBoard.getPosition(2, 4) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	err = theBoard.makeMove(15, 10)
	if err != nil {
		t.Errorf("This should be just fine")
	}
	if theBoard.getPosition(0, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
}

func TestSimpleWin(t *testing.T) {
	theBoard := newBoard(nil)

	// theBoard.printBoard()

	// some one should win
	theBoard.makeMove(0, 5) //x go
	theBoard.makeMove(1, 6) //o go
	// 2
	// theBoard.printBoard()
	theBoard.makeMove(0, 5) //x go
	theBoard.makeMove(1, 6)
	//4
	theBoard.makeMove(0, 5)
	theBoard.makeMove(1, 6)
	// 6
	theBoard.makeMove(0, 5)
	theBoard.makeMove(1, 6)
	// 8
	theBoard.makeMove(0, 5)
	theBoard.makeMove(1, 6)
	// 10

	if theBoard.checkForWin() == '#' {
		t.Errorf("this should be a win")
	}
}

func TestSimpleWin2(t *testing.T) {

	theBoard := newBoard(nil)

	// theBoard.printBoard()

	// some one should win
	theBoard.makeMove(10, 15) //x go
	theBoard.makeMove(11, 16) //o go
	// 2
	// theBoard.printBoard()
	theBoard.makeMove(10, 15) //x go
	theBoard.makeMove(11, 16) //o go
	//4
	theBoard.makeMove(10, 15) //x go
	theBoard.makeMove(11, 16) //o go
	// 6
	theBoard.makeMove(10, 15) //x go
	theBoard.makeMove(11, 16) //o go
	// 8
	theBoard.makeMove(10, 15) //x go
	theBoard.makeMove(11, 16) //o go

	if theBoard.checkForWin() == '#' {
		t.Errorf("this should be a win")
	}
}

func Test12DiagWin(t *testing.T) {

	theBoard := newBoard(nil)

	// theBoard.printBoard()

	// the diagonal 12 win
	theBoard.makeMove(9, 4)
	theBoard.makeMove(7, 2)

	theBoard.makeMove(8, 3)
	theBoard.makeMove(8, 3)

	theBoard.makeMove(7, 2)
	theBoard.makeMove(7, 2)

	theBoard.makeMove(6, 1)
	theBoard.makeMove(7, 2)

	theBoard.makeMove(5, 0)
	theBoard.makeMove(5, 0)
	theBoard.makeMove(5, 0)
	theBoard.makeMove(5, 0)
	theBoard.makeMove(5, 0)

	theBoard.makeMove(6, 1)
	theBoard.makeMove(6, 1)
	theBoard.makeMove(6, 1)

	// theBoard.printBoard()
	if theBoard.checkForWin() != 'X' {
		t.Errorf("this should be a win")
	}
}

//should be illegal
func TestIlegalMove1(t *testing.T) {
	t.Log("this is testIlegalMove1")
	theBoard := newBoard(nil)
	theBoard.makeMove(0, 5)        //x go
	err := theBoard.makeMove(5, 0) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestIlegalMove1r(t *testing.T) {
	theBoard := newBoard(nil)
	theBoard.makeMove(5, 0)        //x go
	err := theBoard.makeMove(0, 5) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestIlegalMove2(t *testing.T) {
	theBoard := newBoard(nil)
	theBoard.makeMove(10, 15)        //x go
	err := theBoard.makeMove(15, 10) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestIlegalMove2r(t *testing.T) {
	theBoard := newBoard(nil)
	theBoard.makeMove(15, 10)        //x go
	err := theBoard.makeMove(10, 15) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestNewBoardWithALayout(t *testing.T) {
	aLayout := [][]rune{
		[]rune{'X', 'O', '#', '#', '#'},
		[]rune{'X', 'O', '#', '#', '#'},
		[]rune{'O', '#', '#', '#', '#'},
		[]rune{'x', 'O', '#', '#', '#'},
		[]rune{'X', 'O', '#', '#', '#'}}
	theBoard := newBoard(aLayout)
	// add board equality
	if !theBoard.equals(newBoard(aLayout)) {
		t.Errorf("The board should be initalized with the layout")
	}
}

func TestForNoWinWhenBothWin(t *testing.T) {
	aLayout := [][]rune{
		[]rune{'X', 'O', '#', '#', '#'},
		[]rune{'X', 'O', '#', '#', '#'},
		[]rune{'O', '#', '#', '#', '#'},
		[]rune{'X', 'O', '#', '#', '#'},
		[]rune{'X', 'O', '#', '#', '#'}}
	theBoard := newBoard(aLayout)
	theBoard.makeMove(17, 12)
	if theBoard.checkForWin() != 'O' {
		t.Errorf("O should win because even though X has five in a row so does O")
	}
}

func TestUperLeftCorner(t *testing.T) {
	theBoard := newBoard(nil)
	err := theBoard.makeMove(0, 15)
	if err != nil {
		t.Errorf("This should not fail!")
	}
	if theBoard.getPosition(0, 4) != 'X' {
		t.Errorf("There really should be an X in this position")
	}
}

func TestUperEdge(t *testing.T) {
	layout := [][]rune{
		[]rune{'O', 'X', '#', 'O', 'O'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'},
		[]rune{'#', '#', '#', '#', '#'}}
	theBoard := newBoard(layout)
	err := theBoard.makeMove(2, 15)
	if err != nil {
		t.Errorf("This should not fail!")
	}
	if theBoard.getPosition(0, 4) != 'X' {
		t.Errorf("There really should be an X in this position")
	}
}

func TestUpperRightCorner(t *testing.T) {
	theBoard := newBoard(nil)
	err := theBoard.makeMove(4, 10)
	if err != nil {
		t.Errorf("This should not fail!")
	}
	if theBoard.getPosition(0, 0) != 'X' {
		t.Errorf("There really should be an X in this position")
	}
}
