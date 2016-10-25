package main

import (
	_ "bytes"
	"fmt"
	"os"
	_ "strconv"
)

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

	// // some one should win
	// tryMove(theBoard, 0) //x go
	// tryMove(theBoard, 1) //o go
	// // 2
	// // theBoard.printBoard()
	// tryMove(theBoard, 0) //x go
	// tryMove(theBoard, 1)
	// //4
	// tryMove(theBoard, 0)
	// tryMove(theBoard, 1)
	// // 6
	// tryMove(theBoard, 0)
	// tryMove(theBoard, 1)
	// // 8
	// tryMove(theBoard, 0)
	// tryMove(theBoard, 1)
	// // 10

	// the diagonal 12 win
	tryMove(theBoard, 9)
	tryMove(theBoard, 7)

	tryMove(theBoard, 8)
	tryMove(theBoard, 8)

	tryMove(theBoard, 7)
	tryMove(theBoard, 7)

	tryMove(theBoard, 6)
	tryMove(theBoard, 7)

	tryMove(theBoard, 5)
	tryMove(theBoard, 5)
	tryMove(theBoard, 5)
	tryMove(theBoard, 5)
	tryMove(theBoard, 5)

	tryMove(theBoard, 6)
	tryMove(theBoard, 6)
	tryMove(theBoard, 6)

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
	fmt.Printf("%c moves %v\n", theBoard.whoseTurn(), x)
	err := theBoard.makeMove(x)
	if err != nil {
		theBoard.printBoard()
		fmt.Printf("%c can't move %v\n", theBoard.whoseTurn(), x)
		panic("ilegal mOVE!")
	}
	if theBoard.checkForWin() {
		fmt.Printf("someone Won\n")
		theBoard.printBoard()
		os.Exit(0)
	}
}
