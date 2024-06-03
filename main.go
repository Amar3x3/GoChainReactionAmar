package main

import (
	"chainreaction/model"
	"fmt"
)

func main() {
	fmt.Println("You are playing chain reaction multiplayer game");

	var newmatrix = model.Matrix{}
	
	
	newmatrix.PlayMove(model.Move{
		Pid: 1,
		Color: "Red",
		Loc: [2]int{0,0},
	})
	newmatrix.PlayMove(model.Move{
		Pid: 2,
		Color: "Green",
		Loc: [2]int{11,0},
	})
	newmatrix.PlayMove(model.Move{
		Pid: 1,
		Color: "Red",
		Loc: [2]int{0,0},
	})
	newmatrix.PlayMove(model.Move{
		Pid: 2,
		Color: "Green",
		Loc: [2]int{11,0},
	})
	newmatrix.PlayMove(model.Move{
		Pid: 1,
		Color: "Red",
		Loc: [2]int{11,0},
	})
	newmatrix.PlayMove(model.Move{
		Pid: 2,
		Color: "Green",
		Loc: [2]int{10,0},
	})
	newmatrix.PlayMove(model.Move{
		Pid: 1,
		Color: "Red",
		Loc: [2]int{11,0},
	})
	newmatrix.Display()
	
	
}