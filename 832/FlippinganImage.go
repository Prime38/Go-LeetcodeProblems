package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i:=0;i<len(A);i++{
		l:=len(A[i])
		for j:=0;j<(l+1)/2;j++{
			t1:=A[i][j]
			t2:=A[i][l-j-1]

			A[i][l-j-1]=t1^1
			A[i][j]=t2^1
			//A[i][j],A[i][l-j-1]=A[i][l-j-1]^1,A[i][j]^1
		}
	}
	return A
}

func main(){
	A:=[][]int{{1,1,0},{1,0,1},{0,0,0}}
	fmt.Println(flipAndInvertImage(A))

}
