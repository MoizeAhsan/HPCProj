package main

import (
	"os"
	"strconv"
    "os/exec"
	"fmt"
	"io/ioutil"
	"time"

)

func MakeBoard(rows, cols int) (board [][]bool) {
	board = make([][]bool, rows+2)
	for i, _ := range board {
		board[i] = make([]bool, cols+2)
	}
	return board
}

func BoardToString(board[][] bool) string{
	var theString string
	for i,x := range board{
		for j,item := range x{
			// fmt.Println(i!=0 && i!= (len(board)-1) && j != 0 && j!=(len(x)-1))
			if i!=0 && i!= (len(board)-1) && j != 0 && j!=(len(x)-1){
				if item != false{
					theString = theString + "*"
				}else{
					theString = theString + " "
				}
			}
		}
		if i!=0 && i!= (len(board)-1){
			theString = theString+"\n"
		}
		// fmt.Printf("\n")
	}	
	return theString[0:len(theString)]
}

func StringToBoard(str []byte,board [][]bool) {
	i:= 1
	j :=0+1
	for _,val := range str {
		if val == '\n'{
			i++
			j = 0+1
			continue
		}
		if val =='1'{//change this if neccessary
				board[i][j] = true
			}else{
				board[i][j] = false
			}
		j++
	}
	return
}

func NextCellState(board [][]bool, row, col int) bool {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i != 0 || j != 0 /*&& row+i != 0 && col+j != 0 && row+i != len(board)-1 && col+j != len(board[1])-1*/{
				if board[row+i][col+j] == true {
					count++
				}
			}
		}
	}
	if board[row][col] == true {
		if count < 2 {
			return false
		} else if count == 2 || count == 3 {
			return true
		} else if count > 3 {
			return false
		}
	} else if board[row][col] == false {
		if count == 3 {
			return true
		}
	}
	return board[row][col]
}

func NextGameState(oldBoard [][]bool, newBoard [][]bool) {
	for i := 1; i < (len(oldBoard)-1); i++ {
		for j := 1; j < (len(oldBoard[i])-1); j++ {
					newBoard[i][j] = NextCellState(oldBoard,i,j)
				}		
	}
}


func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: go run gol.go <filename> <rows> <cols> <iterations>")
		return
	}
	rows, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
    cols, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
    iters, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println(err)
		return
	}
	gofile, _ := ioutil.ReadFile(os.Args[1])
	oldBoard := MakeBoard(rows, cols)
	StringToBoard(gofile, oldBoard)

	newBoard := MakeBoard(rows, cols)
	for i := 0; i < iters; i++ {
		NextGameState(oldBoard, newBoard)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		fmt.Print("\n\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  GENERATION #")
		fmt.Print(i+1)
		fmt.Println("  @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Print(BoardToString(newBoard))
		oldBoard, newBoard = newBoard, oldBoard
		time.Sleep(time.Second / 30)
	}
}
