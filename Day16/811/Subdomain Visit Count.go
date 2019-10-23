package main

import (
	"fmt"
	"strconv"
	"strings"
)
func subdomainVisits(cpdomains []string) []string {
	m:=make(map[string]int)
	ans:=[]string{}
	for i:=0;i<len(cpdomains);i++{
		elems:=[]string{}
		elems=strings.Split(cpdomains[i]," ")
		name:=elems[1]
		count,_:=strconv.Atoi(elems[0])
		domains:=strings.Split(name,".")
		st:=""
		for j:=len(domains)-1;j>=0;j--{
			if j==len(domains)-1{
				st=domains[j]+st
			} else{
				st=domains[j]+"."+st
			}
			m[st]+=count
		}
	}
	for i,v:=range m{

		a:=strconv.Itoa(v)+" "+i
		ans=append(ans, a)
	}
	return ans
}
func main(){
	cpdomains:=[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(cpdomains))
}
