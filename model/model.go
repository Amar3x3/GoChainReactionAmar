package model

import "fmt"

type Matrix struct {
	Board [rows][cols]int
	Color [rows][cols]string
}

type Move struct {
	Pid   int
	Color string
	Loc   [2]int
}

var Colors = []string{"Blue", "Red", "Green", "Yellow", "Purple", "Orange", "Pink", "Black"}

const rows int = 12
const cols int = 6

var players int = 3

func (matrix *Matrix) PlayMove(move Move) {
	if  matrix.Color[move.Loc[0]][move.Loc[1]] ==move.Color || matrix.Color[move.Loc[0]][move.Loc[1]]==""{
	matrix.Board[move.Loc[0]][move.Loc[1]] += 1
	matrix.Color[move.Loc[0]][move.Loc[1]] = move.Color
	matrix.Reaction(move)
	matrix.isWin()
	}
}

func (matrix *Matrix) Display() {
	fmt.Println(matrix.Board,matrix.Color,matrix.isWin())
}

func (matrix *Matrix) Reaction(move Move) {
	flag := false
	if matrix.Board[0][0] > 1 {
		matrix.Board[0][1] += 1
		matrix.Color[0][1] = move.Color
		matrix.Board[1][0] += 1
		matrix.Color[1][0] = move.Color
		matrix.Board[0][0] = 0
		matrix.Color[0][0] = ""
		flag = true
	}
	if matrix.Board[rows-1][cols-1] > 1 {
		matrix.Board[rows-1][cols-2] += 1
		matrix.Color[rows-1][cols-2] = move.Color
		matrix.Board[rows-2][cols-1] += 1
		matrix.Color[rows-2][cols-1] = move.Color
		matrix.Board[rows-1][cols-1] = 0
		matrix.Color[rows-1][cols-1] = ""
		flag = true
	}
	if matrix.Board[0][cols-1] > 1 {
		matrix.Board[0][cols-2] += 1
		matrix.Color[0][cols-2] = move.Color
		matrix.Board[1][cols-1] += 1
		matrix.Color[1][cols-1] = move.Color
		matrix.Board[0][cols-1] = 0
		matrix.Color[0][cols-1] = ""
		flag = true
	}
	if matrix.Board[rows-1][0] > 1 {
		matrix.Board[rows-1][1] += 1
		matrix.Color[rows-1][1] = move.Color
		matrix.Board[rows-2][0] += 1
		matrix.Color[rows-2][0] = move.Color
		matrix.Board[rows-1][0] = 0
		matrix.Color[rows-1][0] = ""
		flag = true
	}

	for i := 1; i < rows; i++ {
		if matrix.Board[i][0] > 2 {
			matrix.Board[i-1][0] += 1
			matrix.Color[i-1][0] = move.Color
			matrix.Board[i+1][0] += 1
			matrix.Color[i+1][0] = move.Color
			matrix.Board[i][1] += 1
			matrix.Color[i][1] = move.Color
			matrix.Board[i][0] = 0
			matrix.Color[i][0] = ""
			flag = true
		}
		if matrix.Board[i][cols-1] > 2 {
			matrix.Board[i-1][cols-1] += 1
			matrix.Color[i-1][cols-1] = move.Color
			matrix.Board[i+1][cols-1] += 1
			matrix.Color[i+1][cols-1] = move.Color
			matrix.Board[i][cols-2] += 1
			matrix.Color[i][cols-2] = move.Color
			matrix.Board[i][cols-1] = 0
			matrix.Color[i][cols-1] = ""
			flag = true
		}
	}
	for i := 1; i < cols; i++ {
		if matrix.Board[0][i] > 2 {
			matrix.Board[0][i-1] += 1
			matrix.Color[0][i-1] = move.Color
			matrix.Board[0][i+1] += 1
			matrix.Color[0][i+1] = move.Color
			matrix.Board[1][i] += 1
			matrix.Color[1][i] = move.Color
			matrix.Board[0][i] = 0
			matrix.Color[0][i] = ""
			flag = true
		}
		if matrix.Board[rows-1][i] > 2 {
			matrix.Board[rows-1][i-1] += 1
			matrix.Color[rows-1][i-1] = move.Color
			matrix.Board[rows-1][i+1] += 1
			matrix.Color[rows-1][i+1] = move.Color
			matrix.Board[rows-2][i] += 1
			matrix.Color[rows-2][i] = move.Color
			matrix.Board[rows-1][i] = 0
			matrix.Color[rows-1][i] = ""
			flag = true
		}
	}
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if matrix.Board[i][j] > 3 {
				matrix.Board[i-1][j] += 1
				matrix.Color[i-1][j] = move.Color
				matrix.Board[i][j-1] += 1
				matrix.Color[i][j-1] = move.Color
				matrix.Board[i+1][j] += 1
				matrix.Color[i+1][j] = move.Color
				matrix.Board[i][j+1] += 1
				matrix.Color[i][j+1] = move.Color
				matrix.Board[i][j] = 0
				matrix.Color[i][j] = ""
				flag = true
			}
		}
	}
	if flag {
		matrix.Reaction(move)
	}
}

func (matrix *Matrix)isWin()bool{
	color :=""
	for i:=0;i<rows;i++{
		for j:=0;j<cols;j++{
			if matrix.Color[i][j]!="" && color==""{
				color = matrix.Color[i][j]
			}
			if matrix.Color[i][j]!=color && matrix.Color[i][j]!="" {
				return false
			}
		}
	}
	return true
}
