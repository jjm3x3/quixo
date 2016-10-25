package main

import (
	"testing"
)

func TestMoves(t *testing.T) {

	theBoard := newBoard()

	theBoard.printBoard()

	theBoard.makeMove(0)
	theBoard.makeMove(0)
	theBoard.makeMove(0)
	theBoard.makeMove(0)

	theBoard.makeMove(0)
	theBoard.makeMove(1)
	theBoard.makeMove(1)
	theBoard.makeMove(1)
	theBoard.makeMove(1)
	theBoard.makeMove(1)
	theBoard.makeMove(5)
	theBoard.makeMove(10)
	theBoard.makeMove(10)
	theBoard.makeMove(11)
	theBoard.makeMove(15)
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
