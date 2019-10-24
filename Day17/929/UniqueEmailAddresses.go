package main

import "fmt"

func numUniqueEmails(emails []string) int {
	ans:=0
	m:=make(map[string]int)
	for i:=0;i<len(emails);i++{
		var str string

		for j:=0;j<len(emails[i]);j++{

			if emails[i][j]=='@'{
				str+=emails[i][j:]
				break

			} else{
				if emails[i][j]=='.' {
					continue
				} else if emails[i][j]=='+'{
					for emails[i][j]!='@'{
						j++
					}
					j--
				} else{
					str+=string(emails[i][j])
				}
			}

		}
		m[str]++
	}

	for _,_=range m{
		ans++
	}
	return ans
}
func main(){
	emails:=[]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
