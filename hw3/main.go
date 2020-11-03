package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var test [][]float64

	test = make([][]float64, 3)
	for r := range test {
		test[r] = make([]float64, 3)
	}

	test[1][2] = 0.6
	test[0][1] = 0.4
	test[0][2] = 1.2

	fmt.Println(BinarizeMatrix(test, 0.5))

}

func InitializeMatrix(row int, col int) [][]int {

	var ret [][]int

	ret = make([][]int, row)

	for r := range ret {
		ret[r] = make([]int, col)
	}

	return ret

}

func BinarizeMatrix(mtx [][]float64, threshold float64) [][]bool {

	var ret [][]bool

	ret = make([][]bool, len(mtx))
	for x := range mtx {
		ret[x] = make([]bool, len(mtx[x]))
	}
	for r := range mtx {
		for c := range mtx[r] {
			if mtx[r][c] >= threshold {
				ret[r][c] = true
			}
		}
	}
	return ret

}

func ConnectFour(mtx [][]bool) bool {

	if getCols(mtx) < 4 || getRows(mtx) < 4 {
		return false
	}

	for r := range mtx {
		for c := range mtx[r] {

			if CheckVertical(mtx, r, c) || CheckDiagonal(mtx, r, c) || CheckHorizontal(mtx, r, c) {
				return true
			}
		}
	}
	return false
}

func CheckVertical(mtx [][]bool, r, c int) bool {
	dy := []int{0, 1, 2, 3}
	dx := []int{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		if !isValid(mtx, r+dy[i], c+dx[i]) {
			return false
		}
		if mtx[r+dy[i]][c+dx[i]] == false {
			return false
		}
	}
	return true
}

func CheckHorizontal(mtx [][]bool, r, c int) bool {
	dy := []int{0, 0, 0, 0}
	dx := []int{0, 1, 2, 3}
	for i := 0; i < 4; i++ {
		if !isValid(mtx, r+dy[i], c+dx[i]) {
			return false
		}
		if mtx[r+dy[i]][c+dx[i]] == false {
			return false
		}
	}
	return true
}

func CheckDiagonal(mtx [][]bool, r, c int) bool {
	dy := []int{0, 1, 2, 3}
	dx := []int{0, 1, 2, 3}
	for i := 0; i < 4; i++ {
		if !isValid(mtx, r+dy[i], c+dx[i]) {
			return false
		}
		if mtx[r+dy[i]][c+dx[i]] == false {
			return false
		}
	}
	return true
}

func getRows(mtx [][]bool) int {
	return len(mtx)
}

func getCols(mtx [][]bool) int {

	if getRows(mtx) == 0 {
		panic("Empty board given")
	}
	return len(mtx[0])
}

func isValid(board [][]bool, r int, c int) bool {
	return r >= 0 && c >= 0 && r < getRows(board) && c < getCols(board)
}

func isValidWalk(row int, col int, width int) bool {
	return row >= 0 && col >= 0 && row < width && col < width

}

func RandomWalk(numSteps int, width int) [][]int {

	ret := make([][]int, numSteps)
	for r := range ret {
		ret[r] = make([]int, 2)
	}
	dy := []int{0, 0, 1, 1, 1, -1, -1, -1}
	dx := []int{1, -1, 1, -1, 1, 1, -1, 1}

	stepChangeIndex := rand.Intn(len(dy))

	ret[0] = []int{0, 0}

	for i := 1; i < len(ret); i++ {
		for !isValidWalk(ret[i-1][0]+dy[stepChangeIndex], ret[i-1][1]+dx[stepChangeIndex], width) {
			stepChangeIndex = rand.Intn(len(dy))
		}
		ret[i] = []int{ret[i-1][0] + dy[stepChangeIndex], ret[i-1][1] + dx[stepChangeIndex]}
	}
	return ret

}
