package test

import "testing"
import (
	"../milestones"
)

func TestFizzBuzz(t *testing.T) {

	var x = []string{"1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"}
	var count int

	for fizzbuzzValue := range milestones.Fizzbuzz(15) {
		if fizzbuzzValue != x[count] {
			t.Errorf("I got %s but I was expecting %s", fizzbuzzValue, x[count])
		}
		count+=1
	}
}


