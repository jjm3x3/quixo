package main

import (
	"testing"
)

func TestMoves(t *testing.T) {

	theBoard := newBoard()

	theBoard.printBoard()

	theBoard.makeMove(0)
	if theBoard.getPosition(4, 0) != 'X' {
		t.Errorf("this should be a plenty fine opening move")
	}
	err := theBoard.makeMove(0)
	if err != nil {
		t.Logf("this is a bad move?!")
	}
	if theBoard.getPosition(4, 0) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(0)
	if theBoard.getPosition(4, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(0)
	if theBoard.getPosition(4, 0) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}

	theBoard.makeMove(0)
	if theBoard.getPosition(4, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1)
	if theBoard.getPosition(4, 1) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1)
	if theBoard.getPosition(4, 1) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1)
	if theBoard.getPosition(4, 1) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1)
	if theBoard.getPosition(4, 1) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(1)
	if theBoard.getPosition(4, 1) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(5)
	if theBoard.getPosition(0, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	err = theBoard.makeMove(10)
	if err == nil {
		t.Errorf("this shouldent be allowed")
	}

	// make expicit turn change for testing
	theBoard.turn = !theBoard.turn

	err = theBoard.makeMove(10)
	if err != nil {
		t.Errorf("This should be just fine")
	}
	if theBoard.getPosition(0, 4) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
	theBoard.makeMove(12)
	if theBoard.getPosition(2, 4) != 'O' {
		t.Errorf("this should be a plenty fine counter move")
	}
	err = theBoard.makeMove(15)
	if err != nil {
		t.Errorf("This should be just fine")
	}
	if theBoard.getPosition(0, 0) != 'X' {
		t.Errorf("this should be a plenty fine counter move")
	}
}

func TestSimpleWin(t *testing.T) {
	theBoard := newBoard()

	// theBoard.printBoard()

	// some one should win
	theBoard.makeMove(0) //x go
	theBoard.makeMove(1) //o go
	// 2
	// theBoard.printBoard()
	theBoard.makeMove(0) //x go
	theBoard.makeMove(1)
	//4
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	// 6
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	// 8
	theBoard.makeMove(0)
	theBoard.makeMove(1)
	// 10

	if !theBoard.checkForWin() {
		t.Errorf("this should be a win")
	}
}

func TestSimpleWin2(t *testing.T) {

	theBoard := newBoard()

	// theBoard.printBoard()

	// some one should win
	theBoard.makeMove(10) //x go
	theBoard.makeMove(11) //o go
	// 2
	// theBoard.printBoard()
	theBoard.makeMove(10) //x go
	theBoard.makeMove(11) //o go
	//4
	theBoard.makeMove(10) //x go
	theBoard.makeMove(11) //o go
	// 6
	theBoard.makeMove(10) //x go
	theBoard.makeMove(11) //o go
	// 8
	theBoard.makeMove(10) //x go
	theBoard.makeMove(11) //o go

	if !theBoard.checkForWin() {
		t.Errorf("this should be a win")
	}
}

func Test12DiagWin(t *testing.T) {

	theBoard := newBoard()

	// theBoard.printBoard()

	// the diagonal 12 win
	theBoard.makeMove(9)
	theBoard.makeMove(7)

	theBoard.makeMove(8)
	theBoard.makeMove(8)

	theBoard.makeMove(7)
	theBoard.makeMove(7)

	theBoard.makeMove(6)
	theBoard.makeMove(7)

	theBoard.makeMove(5)
	theBoard.makeMove(5)
	theBoard.makeMove(5)
	theBoard.makeMove(5)
	theBoard.makeMove(5)

	theBoard.makeMove(6)
	theBoard.makeMove(6)
	theBoard.makeMove(6)

	if !theBoard.checkForWin() {
		t.Errorf("this should be a win")
	}
}

//should be illegal
func TestIlegalMove1(t *testing.T) {
	t.Log("this is testIlegalMove1")
	theBoard := newBoard()
	theBoard.makeMove(0)        //x go
	err := theBoard.makeMove(5) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestIlegalMove1r(t *testing.T) {
	theBoard := newBoard()
	theBoard.makeMove(5)        //x go
	err := theBoard.makeMove(0) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestIlegalMove2(t *testing.T) {
	theBoard := newBoard()
	theBoard.makeMove(10)        //x go
	err := theBoard.makeMove(15) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}

func TestIlegalMove2r(t *testing.T) {
	theBoard := newBoard()
	theBoard.makeMove(15)        //x go
	err := theBoard.makeMove(10) //o go
	if err == nil {
		t.Errorf("the last move is ilegal")
	}
}
