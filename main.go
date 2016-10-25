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
