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

var (
	theMoveList [][]int
)

func main() {

	theBoard := newBoard(nil)
	theBoard.makeMove(0, 5)

	moveList := getMoves()

	nextState := make([]*board, 46)

	fmt.Printf("how long is move list %v\n", len(moveList))
	for i := 0; i < len(moveList); i++ {
		theMove := moveList[i]
		src := theMove[0]
		dest := theMove[1]
		err := theBoard.checkMove(src, dest)
		if err == nil {
			newState := copy(theBoard)
			newState.makeMove(src, dest)
			nextState[i] = newState
		}
	}

	fmt.Printf("show me what you look like:\n %v\n", nextState)

	newNextState := make([][]*board, len(nextState))
	for i := 0; i < len(nextState); i++ {
		startState := nextState[i]
		newNextState[i] = make([]*board, len(moveList))
		for j := 0; j < len(moveList); j++ {
			if startState != nil {
				err := startState.checkMove(moveList[j][0], moveList[j][1])
				if err == nil {
					newState := copy(startState)
					newState.makeMove(moveList[j][0], moveList[j][1])
					newNextState[i][j] = newState
				}
			}
		}
	}
	fmt.Printf("lets see theseNext moves:\n %v\n", newNextState)

	possibleMoves := howManyMoves(theBoard)
	fmt.Printf("how many moves can I make?: %v\n", possibleMoves)

	// getMoves()

	// playGame()

}

func howManyMoves(board *board) int {
	possibleMoves := 0
	for i := 0; i < len(theMoveList); i++ {
		if board.checkMove(theMoveList[i][0], theMoveList[i][1]) == nil {
			possibleMoves++
		}
	}
	return possibleMoves
}

func getMoves() [][]int {

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

	result := make([][]int, len(moves))

	for i := 0; i < len(moves); i++ {
		theMove := moves[i]
		src, err := strconv.Atoi(theMove[0])
		if err != nil {
			log.Printf("problem casting src is not a number: %v\n", err)
		}
		dest, err := strconv.Atoi(theMove[1])
		if err != nil {
			log.Printf("problem casting src is not a number: %v\n", err)
		}
		result[i] = []int{src, dest}
	}

	// for i := 0; i < len(moves); i++ {
	// 	fmt.Printf("I see a move:  %v\n", moves[i])
	// }

	theMoveList = result // not sure which is preferable
	return result
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
