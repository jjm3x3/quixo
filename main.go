package main

import (
	"bufio"
	_ "bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
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
				move := findNextMove(theBoard, BASIC_AI)
				theBoard.makeMove(move[0], move[1])
				checkForWin(theBoard)
			}
		}
	} else {
		playGame()
	}

}

func findNextMove(theBoard *board, aiKind int) []int {

	if aiKind == NEURAL_NETWORK {
		return neuralNetwork(theBoard)
	} else {
		return basicAI(theBoard)
	}

}

func neuralNetwork(theBoard *board) []int {
	// numberOfMoves := howManyMoves(theBoard)
	// numberOfPeices := howManyPeices(theBoard, theBoard.whoseTurn())

	nextStates := getNextStates(theBoard)

	// fmt.Printf("show me what you look like:\n %v\n", nextStates)

	var (
		bestMove  []int
		bestScore = float64(50)
	)
	for i := 0; i < len(nextStates); i++ {
		if nextStates[i] != nil {
			newMoves := howManyMoves(nextStates[i])
			numPeices := howManyPeices(nextStates[i], theBoard.whoseTurn())
			outcome := determineOutcome(newMoves, numPeices)
			log.Printf("yeilds result: %v vs current best: %v\n", outcome, bestScore)
			if outcome < bestScore {
				bestScore = outcome
				bestMove = theMoveList[i]
			}

		}
	}

	return bestMove

}

type neuron struct {
	w1,
	w2 float64
}

func (n *neuron) compute(x, y float64) float64 {
	return float64(x)*n.w1 + float64(y)*n.w2
}

func determineOutcome(x, y int) float64 {
	log.Printf("%v , %v\n", x, y)
	now := time.Now()
	rand.Seed(int64(now.Nanosecond()))
	// w1 := rand.Intn(45)
	// w2 := rand.Intn(45)

	// return x*w1 + y*w2

	fx := float64(x)
	fy := float64(y)
	firstN := &neuron{0.5, 0.5}
	firstN1 := &neuron{0.1, 0.2}
	finalN := &neuron{0.7, 0.36}

	return finalN.compute(firstN.compute(fx, fy), firstN1.compute(fx, fy))

}

func basicAI(theBoard *board) []int {
	nextStates := getNextStates(theBoard)

	// fmt.Printf("show me what you look like:\n %v\n", nextStates)

	var (
		bestMove   []int
		mostMoves  int
		mostPeices int
	)
	for i := 0; i < len(nextStates); i++ {
		if nextStates[i] != nil {
			newMoves := howManyMoves(nextStates[i])
			// if newMoves > mostMoves {
			// 	mostMoves = newMoves
			// }

			numPeices := howManyPeices(nextStates[i], theBoard.whoseTurn())
			if numPeices > mostPeices {
				mostPeices = numPeices
				bestMove = theMoveList[i]
			} else if numPeices == mostPeices && newMoves >= mostMoves {
				bestMove = theMoveList[i]
			}

		}
	}

	return bestMove
}

func getNextStates(startState *board) []*board {

	nextStates := make([]*board, len(theMoveList))

	for i := 0; i < len(theMoveList); i++ {
		theMove := theMoveList[i]
		src := theMove[0]
		dest := theMove[1]
		err := startState.checkMove(src, dest)
		if err == nil {
			newState := copy(startState)
			newState.makeMove(src, dest)
			nextStates[i] = newState
		}
	}

	return nextStates
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

func howManyPeices(board *board, player rune) int {
	result := 0
	for i := 0; i < len(board.grid); i++ {
		for j := 0; j < len(board.grid[i]); j++ {
			if board.getPosition(i, j) == player {
				result++
			}
		}
	}
	return result
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
			move := findNextMove(theBoard, BASIC_AI)
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
