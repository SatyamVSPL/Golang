package main

func Factorial(i int) int{
	// var i int 
	fact:=1
	// fmt.Print("Enter the Number : ")
	// fmt.Scanln(&i)
	for num:=i ; num>0 ; num--{
		fact*= num
	}
	return fact
}