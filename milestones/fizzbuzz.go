package milestones

import (
	"strconv"
)

func Fizzbuzz(value int) chan string {

	//buffered channel that holds the string values
	channel := make(chan string, value)

	//compute 100 values
	for i := 1; i <= value; i++ {
		result := ""
		if i%3 == 0 {
			result = "fizz"
		}
		if i%5 == 0 {
			result += "buzz"
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		channel <- result
	}

	//close the channel to terminate the range loop
	close(channel)

	return channel
}
