package main

import (
	"fmt"
	"../milestones"

)


func main() {
	var x = []string{"1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"}
	var count int
	for fizzbuzzValue := range milestones.Fizzbuzz(15) {
		if fizzbuzzValue != x[count] {
			fmt.Println(x[count])
			fmt.Println(fizzbuzzValue)
			fmt.Println("ERROR")
		}
		count+=1
	}
}
