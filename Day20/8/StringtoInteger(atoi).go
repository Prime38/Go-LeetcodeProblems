package main

import (
	"fmt"
	"math"
	"strings"
)
func charToInt(r uint8) int {
	return int(r)-48
}
func inrange(r uint8)bool{
	return charToInt(r)>=0&& charToInt(r)<=9
}
func myAtoi(str string) int {
	ans:=0
	ma:=int(math.Pow(2,31)-1)
	mini:= - int(math.Pow(2,31))
	str=strings.Trim(str," ")
	if len(str)==0{
		return 0
	} else if str[0]=='-'{
		if len(str)==1{
			return 0
		} else if str[1]==' '{
			return 0
		} else if inrange(str[1]){
			for i:=1; i<len(str);i++{
				if inrange(str[i]){
					ans=ans*10-charToInt(str[i])
					if ans< mini{
						return mini
					}
				} else{
					break
				}

			}
		} else{
			return 0
		}
	} else if str[0]=='+'|| inrange(str[0]){
		i:=0
		if str[0]=='+'{
			i++
		}
		for ;i<len(str);i++{
			  if inrange(str[i]){
				ans=ans*10+charToInt(str[i])
				if ans> ma{
					return ma
				}
			} else{
				break
			}
		}
	}
	return ans
}
func main(){

	str:="+3+123213"
	fmt.Println(myAtoi(str))




}
