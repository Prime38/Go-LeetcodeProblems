package main

func defangIPaddr(address string) string {
	var ans string
	for i:=0;i< len(address);i++ {
		if address[i]=='.' {
			ans+="[.]"
		} else {
			ans+=string(address[i])
		}
	}
	return ans
}


func main(){
	address:="255.100.50.0"
	print(defangIPaddr(address))

}