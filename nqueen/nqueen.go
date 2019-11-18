package main

import (
	"fmt"
)

const N int = 4

var Board [N][N]int = [N][N]int{{}} //初始化棋盘
var Result int = 0

func setBlock(row int, col int, val int) {
	//设置不可放置宫格
	for r, c := row, col; r < N && c >= 0; {
		//设置左下宫格
		Board[r][c] += val
		c--
		r++
	}
	for r := row; r < N; r++ {
		//设置正下方宫格
		Board[r][col] += val
	}
	for r, c := row, col; r < N && c < N; {
		//设置右下方宫格
		Board[r][c] += val
		c++
		r++
	}
}

func recursive(row int, col int) {
	if row == N {
		Result++
		return
	}
	for ; col < N; col++ {
		if Board[row][col] == 0 {
			//设置不可访问宫格
			setBlock(row, col, -1)
			recursive(row+1, 0)
			//还原上一次设置
			setBlock(row, col, 1)
		}
	}
}

func main() {
	recursive(0, 0)
	fmt.Println(Result)
}
