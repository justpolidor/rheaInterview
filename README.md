# Rhea quizzes

## Milestone 1

_Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz"._

Under the Milestones package, we can find the **fizzbuzz.go** source file. Here we have only one function that will perform the computation over 100 numbers. 

`func Fizzbuzz(value int) chan string` 

The function accepts the integer and it returns a channel which we can retrieve the result; in this case we return a string because we can have three kind of outputs:
1. A number
2. Fizz/Buzz
3. Fizzbuzz

Inside the function we initialize the channel and we perform the computation and we do the following:
* if the number is divisible by 3, so append to the string the word "fizz";
* if the number is divisible by 5, so appent to the string the word "buzz";
* if we had not append anything to the string, so we return the number

notice that with the append we did also the part _For numbers which are multiples of both three and five print "FizzBuzz"_

after we calculated all the results (and we streamed them over the channel), we should close the channel to avoid deadlock and to stop the range iteration over it. 

The unit test that I have written checks the first 15 resuls. If the array with the results are the same of the results from the channel, the test pass. 

## Milestone 2

_If the numbers 1 to 5 are written out in words: "one", "two", "three", "four", "five", then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total to spell out the words._
 
_If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?_

Under the milestones package we can find the **charscount.go** source file. Here we have a function with a switch in it that will perform all the logic; Here I use recursion to achieve the result; 
The `computeNumber(number int, and bool) int` function has two parameters:
1. The number to compute
2. A boolean that we check so we can use the form "one-thousand-**and**-four (if _and_ is true) or "one-thousand-four" (if _and_ is false)
The basic principle is to divide the number into dozens and units (using module and divisions) and check them with the switch/case structures.
After we computed all the 999 numbers, we add the length of _ONETHOUSAND_ (beacuse it is the last number, we just add it without going again under the `computeNumber(number int, and bool) int` function.

Note that the test file checks both _and_ cases (with _and_ true and with _and_ false)
