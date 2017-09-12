package milestones

import (
	"strconv"
)

func Fizzbuzz(limit int) chan string {

	//buffered channel that holds the string values
	resultStream := make(chan string, limit)

	go func() {
		//compute 100 values
		for i := 1; i <= limit; i++ {
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
			resultStream <- result
		}
		//close the channel to terminate the range loop
		close(resultStream)
	}()

	return resultStream
}
