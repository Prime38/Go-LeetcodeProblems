package main

import (
	"bufio"
	"os"
)

func main() {
	cin := bufio.NewReader(os.Stdin)

	n, _ := cin.ReadString('\n')
	len:=len(n)
	for i:=0; i<len; i++ {
		println(string (n[i]))
	}

}