package main

import "fmt"

func diStringMatch(S string) []int {
	I,D:=0,len(S)
	ans:=[]int{}
	if S[0]=='I'{
		ans=append(ans,I)
		I++
	} else{
		ans=append(ans,D)
		D--
	}
	for i:=1;i<len(S);i++{
		if S[i-1]=='I'{
			if S[i]=='D'{
				ans=append(ans,D)
				D--
			} else{
				ans=append(ans,I)
				I++
			}
		} else{
			if S[i]=='I'{
				ans=append(ans,I)
				I++
			} else{
				ans=append(ans,D)
				D--
			}
		}
	}
	if S[len(S)-1]=='I'{
		ans=append(ans,I)
		I++
	} else{
		ans=append(ans,D)
		D--
	}
	return ans
}

func main(){
	S:="DDI"
	fmt.Println(diStringMatch(S))

}
