package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {

	theBoard := newBoard(nil)

	for {
		theBoard.printBoard()

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("select a move\n")
		text, _ := reader.ReadString('\n')
		move, err := strconv.Atoi(text[0 : len(text)-1])
		if err != nil {
			fmt.Printf("There was an error parsing the move: %v\n", err)
		}
		fmt.Printf("select a destination\n")
		text, _ = reader.ReadString('\n')
		dest, err := strconv.Atoi(text[0 : len(text)-1])
		if err != nil {
			fmt.Printf("There was an error parsing the move: %v\n", err)
		}
		// fmt.Printf("what is the number %d\n", move)
		tryMove(theBoard, move, dest)

	}

}

func tryMove(theBoard *board, x, y int) {
	// fmt.Printf("whose turn? %v\n", theBoard.turn)
	fmt.Printf("%c moves %v\n", theBoard.whoseTurn(), x)
	err := theBoard.makeMove(x, y)
	if err != nil {
		theBoard.printBoard()
		fmt.Printf("%c can't move %v\n", theBoard.whoseTurn(), x)
		panic("ilegal mOVE!")
	}
	if theBoard.checkForWin() != '#' {
		fmt.Printf("someone Won\n")
		theBoard.printBoard()
		os.Exit(0)
	}
}
