package main

import "fmt"

func judgeCircle(moves string) bool {
	u:=0
	l:=0
	for i:=0;i<len(moves);i++{
		if moves[i]=='U'{
			u+=1
		} else if moves[i]=='D'{
			u-=1
		} else if moves[i]=='L'{
			l+=1
		} else if moves[i]=='R'{
			l-=1
		}
	}
	return (u==0 && l==0)
}

func main(){
	moves:="LL"
	fmt.Println(judgeCircle(moves))

}
