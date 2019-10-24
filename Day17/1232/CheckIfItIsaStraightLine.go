package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	x1:=coordinates[0][0]
	y1:=coordinates[0][1]
	x2:=coordinates[1][0]
	y2:=coordinates[1][1]
	slope:=float32((y2-y1)/(x2-x1))
	for i:=2;i<len(coordinates);i++{
		x2,y2=coordinates[i][0],coordinates[i][1]
		if slope!=float32((y2-y1)/(x2-x1)){
			return false
		}
	}
 return true
}

func main(){
	co:=[][]int{{1,2},{2,3},{3,4},{5,4},{5,6}}
	fmt.Println(checkStraightLine(co))
}
