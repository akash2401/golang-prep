package main

import "fmt"

func reverseNumber(num int) int {

	reverseNum := 0
	remainder := 0
	for num > 0 {
		remainder = num % 10
		reverseNum = (reverseNum * 10) + remainder
		num = num / 10
	}
	return reverseNum
}

func main() {

	fmt.Println("Reverse of Number is: ", reverseNumber(152351))
}
