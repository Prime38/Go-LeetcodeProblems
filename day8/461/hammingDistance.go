package main

import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) int {
	ans:=0
	a:=int64(x^y)
	str:=strconv.FormatInt(a, 2)
	for i:=0;i<len(str);i++{
		if str[i]=='1'{
			ans++
		}

	}
	return ans
}

func main()  {
	x:=1
	y:=4
	fmt.Println(hammingDistance(x,y))
	
}
