package main

import (
	"bufio"
	_ "bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	theBoard := newBoard(nil)
	theBoard.makeMove(0, 5)

	moveList := getMoves()

	for i := 0; i < len(moveList); i++ {
		theMove := moveList[i]
		src, err := strconv.Atoi(theMove[0])
		if err != nil {
			log.Printf("problem casting src is not a number: %v\n", err)
		}
		dest, err := strconv.Atoi(theMove[1])
		if err != nil {
			log.Printf("problem casting src is not a number: %v\n", err)
		}
		err = theBoard.makeMove(src, dest)
		if err != nil {

		}
	}

	// getMoves()

	// playGame()

}

func getMoves() [][]string {

	file, err := os.Open("posibleMoves.csv")
	defer file.Close()
	if err != nil {
		log.Printf("Error opein move list: %v\n", err)
		panic("I cannot go on!")
	}

	r := csv.NewReader(file)
	moves, err := r.ReadAll()
	if err != nil {
		log.Printf("Error opein move list: %v\n", err)
	}

	for i := 0; i < len(moves); i++ {
		fmt.Printf("I see a move:  %v\n", moves[i])
	}

	return moves
}

func playGame() {

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
