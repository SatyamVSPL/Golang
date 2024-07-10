// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// 	// "time"
// )

// func printStatus(s string, c chan string) {
// 	_, err := http.Get(s)
// 	if err != nil {
// 		fmt.Println("Some isse with :", s)
// 		c <- s
// 		return
// 	}
// 	fmt.Println(s, " Working great")
// 	c <- s

// }

// func main() {
// 	links := []string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://stackoverflow.com",
// 		"http://localhost:3000.com",
// 	}

// 	C := make(chan string)

// 	for _, link := range links {
// 		go printStatus(link, C)
// 		// time.Sleep(2 * time.Second)

// 	}
// 	// fmt.Println(<-C)
// 	// fmt.Println(<-C)
// 	// fmt.Println(<-C)
// 	// fmt.Println(<-C)
// 	// fmt.Println(len(links))

// 	// for range links{
// 	// 	fmt.Println(<-C)
// 	// }

// 	// for link :=range C {
// 	// 	go printStatus(link , C)
// 	// }

// 	// wait for 5 sec
// 	for link :=range C {
// 		go func (l string)  {
// 			time.Sleep(5* time.Second)
// 			printStatus(link , C)
// 		}(link)
		
// 	}
// }


package main
 
import "fmt"
 
func main() {
 c := make(chan string)
 
 for i := 0; i < 4; i++ {
     go printString("Hello there!", c)
 }
 
 for {
     fmt.Println(<- c)
 }
}
 
func printString(s string, c chan string) {
 fmt.Println(s)
 c <- "Done printing." 
}