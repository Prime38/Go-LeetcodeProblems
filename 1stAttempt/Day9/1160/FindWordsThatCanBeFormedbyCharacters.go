package main

import "fmt"

func countCharacters(words []string, chars string) int {
	// using map -> bad solution
	//m:=make(map[byte]int)
	//for i:=0;i<len(chars);i++{
	//	m[chars[i]]++
	//}
	//ans:=0
	//for i:=0;i<len(words);i++{
	//	m2:=map[byte]int{}
	//	for k,v:=range m{
	//		m2[k]=v
	//	}
	//	word:=words[i]
	//	exists:=true
	//	for j:=0;j<len(word);j++{
	//		if val,ok:=m2[word[j]];ok{
	//			if val<=0{
	//				exists=false
	//				break
	//			} else{
	//				m2[word[j]]--
	//			}
	//		} else{
	//			exists=false
	//			break
	//		}
	//
	//	}
	//	if exists==true{
	//		ans+=len(word)
	//	}
	//}
	//return ans

	//using array
	m:=make([]int,28)
	for i:=0;i<len(chars);i++{
		m[int(chars[i])-97]++
	}
	ans:=0
	for i:=0;i<len(words);i++{
		m2:=make([]int,len(m))
		copy(m2,m)
		word:=words[i]
		exists:=true
		for j:=0;j<len(word);j++{
			if m2[int(word[j])-97]>0{
				m2[int(word[j])-97]--
			} else{
				exists=false
				break
			}
		}
		if exists==true{
			ans+=len(word)
		}
	}
	return ans
}

func main(){
	words:=[]string{"hello","world","leetcode"}
	chars:="welldonehoneyr"
	fmt.Println(countCharacters(words,chars))
}
