package main

import "fmt"

func sumNumber(num... int)int{
	sum := 0
	for _,i:= range num {
		sum += i
	}
	return sum
}
func main() {
	num :=[]int{1,2,3,4,5}
	sum := sumNumber(num...)
	fmt.Println(sum)
}

