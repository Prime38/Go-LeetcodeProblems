package main

import "fmt"

func minDeletionSize(A []string) int {
	ans:=0
	for i:=0;i<len(A[0]);i++{
		for j:=0;j<len(A)-1;j++{
			if (int(A[j][i])-97)>(int(A[j+1][i])-97){
				ans++
				break
			}
		}
	}
	return ans
}
func main(){
	A:=[]string{"zyx","wvu","tsr"}
	fmt.Println(minDeletionSize(A))

}
