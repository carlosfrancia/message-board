package main

import "fmt"

//       See types.go for definitions of Board, Move, Moves and more!

func SolveLetterboard(board Board, word string) Moves {
	// moves := []Move{}
	moves := make([]Move, 0)
	//	var letter rune
	for _, char := range word {

		printBoard(board)
		fmt.Printf("Seeking character %v\n", string(char))
		println("LENGTH: ", len(board)/2)
		if getPosition(board, string(char)) <= len(board)/2 {
			moveLeft(&board, char, &moves)
			println("LEFT")
		} else {
			println("RIGHT")
			moveRight(&board, char, &moves)
		}
	}
	return moves
}

func moveRight(board *Board, char rune, moves *[]Move) {
	var pos int

	for pos = len(*board) - 1; pos > 0; pos-- {
		if string((*board)[pos]) == string(char) {
			mv := Move{Right, (*board)[pos]}
			*moves = append(*moves, mv)
			break
		} else {
			mv := Move{Right, -1}
			*moves = append(*moves, mv)
		}
	}
	for i := len(*board) - 1; i > pos; i-- {
		*board = append([]rune{(*board)[len(*board)-1]}, *board...)[:len(*board)-1]
		*board = (*board)[:len(*board)-1]
	}
	*board = (*board)[:len(*board)-1]
}

func moveLeft(board *Board, char rune, moves *[]Move) {
	var pos int
	var letter rune

	for pos, letter = range *board {
		if string(letter) == string(char) {
			mv := Move{Left, letter}
			*moves = append(*moves, mv)
			break
		} else {
			mv := Move{Left, -1}
			*moves = append(*moves, mv)
		}
	}
	for i := 0; i < pos; i++ {
		*board = append(*board, (*board)[0])
		*board = (*board)[1:]
	}
	*board = (*board)[1:]
}

func getPosition(board Board, targetLetter string) int {
	var pos int
	var letter rune
	for pos, letter = range board {
		if string(letter) == targetLetter {
			println("POS: ", pos+1)
			return pos + 1
		}
	}

	return pos
}

func printBoard(board Board) {

	var letter rune
	for _, letter = range board {
		print(string(letter), " ")
	}
	println()

}

func main() {
	// Starting
	print("starting: ")
	println(SolveLetterboard([]rune{'a', 'z', 'c', 't', 'v', 'a'}, "cat"))

}
