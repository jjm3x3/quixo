package main

import (
	"bufio"
	_ "bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	BASIC_AI       = 1
	NEURAL_NETWORK = 2
)

var (
	theMoveList [][]int
	showBoard   bool
)

func main() {
	var numPlayers int
	flag.IntVar(&numPlayers, "players", 0, "This is how many players you want to play the game")
	flag.BoolVar(&showBoard, "show", false, "Toggles wheather the board is printed each time")
	flag.Parse()

	getMoves()

	if numPlayers == 0 {
		playBots()
	} else if numPlayers == 1 {
		theBoard := newBoard(nil)
		for {

			if theBoard.turn == true {
				promptForMove(theBoard)
			} else {
				move := findNextMove(theBoard, NEURAL_NETWORK) //BASIC_AI)
				theBoard.makeMove(move[0], move[1])
				checkForWin(theBoard)
			}
		}
	} else {
		playGame()
	}

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

func playBots() {
	theBoard := newBoard(nil)

	now := time.Now()
	timeString := now.Format(time.RFC3339)
	timeString = timeString[:len(timeString)-6]
	// log.Printf("what time is it: %v", timeString)

	file1, err := os.Create("./games/player1game" + timeString + ".moves")
	if err != nil {
		panic(fmt.Sprintf("how can I track my reslts with: %v\n", err))
	}
	defer file1.Close()

	file2, err := os.Create("./games/player2game" + timeString + ".moves")
	if err != nil {
		panic(fmt.Sprintf("how can I track my reslts with: %v\n", err))
	}
	defer file2.Close()

	for {
		if theBoard.turn {
			move := findNextMove(theBoard, NEURAL_NETWORK)
			theBoard.makeMove(move[0], move[1])
			file1.WriteString(strconv.Itoa(move[0]) + "," + strconv.Itoa(move[1]) + "\n")
		} else {
			move := findNextMove(theBoard, NEURAL_NETWORK) //BASIC_AI)
			theBoard.makeMove(move[0], move[1])
			file2.WriteString(strconv.Itoa(move[0]) + "," + strconv.Itoa(move[1]) + "\n")
		}

		checkForWin(theBoard)
	}
}
func playGame() {

	theBoard := newBoard(nil)

	for {
		promptForMove(theBoard)
		checkForWin(theBoard)
	}
}

func promptForMove(theBoard *board) {

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

func tryMove(theBoard *board, x, y int) {
	// fmt.Printf("whose turn? %v\n", theBoard.turn)
	fmt.Printf("%c moves %v\n", theBoard.whoseTurn(), x)
	err := theBoard.makeMove(x, y)
	if err != nil {
		theBoard.printBoard()
		fmt.Printf("%c can't move %v\n", theBoard.whoseTurn(), x)
		panic("ilegal mOVE!")
	}
}

func checkForWin(board *board) {
	if board.checkForWin() != '#' {
		// need to do the oposite because checkForWin only happen
		// after a move is made implying that whose ever turn
		// turn it is is likely the looser
		if board.whoseTurn() == 'X' {
			fmt.Printf("O won\n")
		} else {
			fmt.Printf("X won\n")
		}
		board.printBoard()
		os.Exit(0)
	}
	if showBoard {
		board.printBoard()
	}

}
