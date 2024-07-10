// package main

// import "fmt"

// func main() {
// 	index:= 1
// 	slice := []int{100,200,30,20,20}
// 	// slice= append(slice,0,100)
// 	// slice[0]
// 	slice = append(slice[:index],slice[index:]...,)
// 	fmt.Println(slice)

// }

package main

import "fmt"

func main() {
	slice := []int{100, 200, 30, 20, 20}
	slice2:= []int{777} // Assign a new integer value directly to slice[0]


	slice = append(slice2, slice... )
	// fmt.Println(slice2)
	fmt.Println(slice) // Print the modified slice
}
