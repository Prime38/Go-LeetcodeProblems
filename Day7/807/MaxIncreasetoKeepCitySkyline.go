package main

import (
	"fmt"
)
func min(x int ,y int)int{
	if x>y{
		return y
	} else {
		return x
	}
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	ans:=0
	maxcol:=[]int{}
	maxrow:=[]int{}
	for i:=0;i<len(grid);i++{
		mr:=0
		mc:=0
		for j:=0;j<len(grid[i]) ;j++  {
			if grid[i][j]>mr{
				mr=grid[i][j]
			}
			if grid[j][i]>mc{
				mc=grid[j][i]
			}
		}
		maxrow=append(maxrow,mr)
		maxcol=append(maxcol,mc)
	}
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]) ;j++  {
			ans+=min(maxrow[i],maxcol[j])-grid[i][j]
		}
	}
	return ans
}

func main(){
	grid:=[][]int{{3,0,8,4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}
