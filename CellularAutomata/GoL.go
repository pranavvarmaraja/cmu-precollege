package main

type GameBoard [][]bool

func PlayGol(initialBoard GameBoard, numGens int) []GameBoard {

	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1])
	}
	return boards
}

func UpdateBoard(board GameBoard) GameBoard {

	numRows := getRows(board)
	numCols := getCols(board)
	newBoard := InitializeBoard(numRows, numCols)

	for r := range board {
		for c := range board[r] {
			newBoard[r][c] = UpdateCell(board, r, c)
		}
	}

	return newBoard
}

func getRows(board GameBoard) int {
	return len(board)
}

func getCols(board GameBoard) int {

	if getRows(board) == 0 {
		panic("Empty board given")
	}
	return len(board[0])
}

func InitializeBoard(rows int, cols int) GameBoard {

	var board GameBoard
	board = make(GameBoard, rows)

	for r := range board {
		board[r] = make([]bool, cols)
	}
	return board
}

func UpdateCell(board GameBoard, row int, col int) bool {

	numNeighbors := countLiveNeighbors(board, row, col)

	if board[row][col] == true {

		if numNeighbors == 2 || numNeighbors == 3 {
			return true
		}
		return false
	} else {
		if numNeighbors == 3 {
			return true
		}
		return false
	}

}

func countLiveNeighbors(board GameBoard, r int, c int) int {

	count := 0
	dy := []int{1, 1, 1, 0, 0, -1, -1, -1}
	dx := []int{0, 1, -1, 1, -1, 0, 1, -1}

	for i := 0; i < 8; i++ {
		if isValid(board, r+dy[i], c+dx[i]) && board[r+dy[i]][c+dx[i]] {
			count++

		}
	}
	return count

}

func isValid(board GameBoard, r int, c int) bool {
	return r >= 0 && c >= 0 && r < getRows(board) && c < getCols(board)
}
