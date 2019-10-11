package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	morse:=[]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	m:=make(map[string]bool)
	l1:=len(words)
	for i:=0;i<l1;i++{
		l2:=len(words[i])
		word:=""
		for j:=0;j<l2;j++ {
			word=word+morse[int(words[i][j])-97]

		}
		fmt.Println(word)
		m[word]=true
	}

	return len(m)
}

func main(){
	words:=[]string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
	fmt.Println('B'-'A')


}
