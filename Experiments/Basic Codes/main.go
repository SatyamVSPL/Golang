package main

import "fmt"

// import "fmt"

// import "fmt"

func main() {
	//even odd
	// fmt.Println(evenodd(4))

	// fact
	// fmt.Println("Factorial :",Factorial(6))

	// pattern
	// fmt.Println(starpattren(4))

	// Prime Number
	// fmt.Println(prime(60))

	// rev str
	// fmt.Println(revStr("Satyam"))

	// palindrome
	// 	fmt.Println(plaindromeStr("BoB"))
	// 	fmt.Println(plaindromeStr("boB"))
	// 	fmt.Println(plaindromeNum(123))
	// fmt.Println(plaindromeNum(12121))

	// slice := []int{1,2,3,4}
	// fmt.Println(slice[1:3])
	// var newSlice[] int

	// fmt.Println(newSlice)

	// newSlice = append(newSlice,1,2,2,3,404,4)

	// fmt.Println(newSlice)

	// index := 4 

	// newSlice = append(newSlice[:index],newSlice[index+1:]...)
	// fmt.Println(newSlice)
	// fmt.Println(newSlice[3:4])


		// name := "Satyam"
		// for index,char := range name{
		// 	fmt.Println("Index of char :",char,"is",index)
		// } 
		// fmt.Println(string(name[4]))

		

	var a number = 10
	// var b number = 20
	// fmt.Println("Value of a :",a,"\nvalue of b :",b)
	// swap(&a ,&b )
	// fmt.Println("Value of a :",a,"\nvalue of b :",b)
	fmt.Println("Value of a :",a)
	a.square()
	fmt.Println("Value of a :",a)

	
}
// func swap(i,j *int) int{
// 	temp := *i
// 	*i = *j
// 	*j = temp 
// 	return 0
	
// }

type number int

func (i *number) square() {
	*i = (*i)*(*i)
}