// sliceToChannle -> ek ek value channel mein dalta jayega
// sq -> channel se read karke square karta jayega

package main

import "fmt"

func sliceToChannle(n []int) <-chan int{
	channel := make(chan int)
	go func(){
		for _,num := range n{
			channel <- num
		}
		close(channel)
	}()
	
	return channel
}	

func sq(num <-chan int) <-chan int{
	channel := make(chan int)
	go func ()  {
		for v := range num{
			channel<- v*v
		}
		close(channel)
	}()
	return channel
}

func main() {
	numSlice := []int {1,2,3,4,5} 
 	slice := sliceToChannle(numSlice)
	result := sq(slice)
	for n := range result{
		fmt.Println(n)
	}
}