package main

import (
	"log"
	"math/rand"
	"time"
)

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

func newNeuron(w1, w2) *neuron {
	return &neuron{w1, w2}
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
	firstN := newNeuron(0.5, 0.5)
	firstN1 := newNeuron(0.1, 0.2)
	finalN := newNeuron(0.7, 0.36)

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
