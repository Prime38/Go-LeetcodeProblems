package main

import "fmt"

func sortedSquares(A []int) []int {
	ans:=[]int{}

     if A[len(A)-1]<0{
		for i:=len(A)-1;i>=0;i--{
			ans=append(ans,A[i]*A[i])
		}
		return ans
	}
	left:=0
	right:=0
	for i:=0;i<len(A);i++{
		if A[i]>=0{
			left=i-1
			right=i
			break
		}
	}
	//rs,ls:=0,0
	//for len(ans)+1!=len(A){
	//	if right<len(A){
	//		rs=A[right]*A[right]
	//	}
	//	if left>=0{
	//		ls=A[left]*A[left]
	//	}
	//	if rs<ls{
	//		ans=append(ans,rs)
	//		right++
	//	} else{
	//		ans=append(ans,ls)
	//		left--
	//	}
	//}
	//if ls>rs{
	//	ans=append(ans,ls)
	//} else{
	//	ans=append(ans,rs)
	//}

	for {
        fmt.Println(left,right)
        if(left < 0 ){
			ans = append(ans, A[right] * A[right])
			right++
		} else if right >= len(A){
			ans = append(ans, A[left] * A[left])
				left--
		} else if ( A[left] * A[left]  < A[right] * A[right] ) {
			ans = append(ans, A[left] * A[left])
				left--


		} else {
			ans = append(ans, A[right] * A[right])

				right++

		}
		if (len(ans) == len(A) ){
			break
		}

	}

	return ans
	//if A[0]>=0{
	//	for i:=0;i<len(A);i++{
	//		A[i]=A[i]*A[i]
	//	}
	//	return A
	//} else
}

func main() {
	A:=[]int{1,2,3,4,5}        // 25, 16,9, || 4,4,16,25
	fmt.Println(sortedSquares(A))     // 4,4,16,16,25,25

}