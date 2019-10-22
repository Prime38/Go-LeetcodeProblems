package main

import (
	"fmt"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	//m:=make(map[int]int)
	for i:=0;i<len(cpdomains);i++{
		elems:=[]string{}
		elems=strings.Split(cpdomains[i]," ")
		name:=elems[1]
		count:=elems[0]
		fmt.Println(count,name)



	}
	return nil
}


func main(){
	cpdomains:=[]string{"9001 discuss.leetcode.com"}
	fmt.Println(subdomainVisits(cpdomains))

}
