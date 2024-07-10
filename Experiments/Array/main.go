package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello");

	// Empty array of size 3
	var Array = [3] string {}
	fmt.Print("This is an empty array ",Array )
	fmt.Println(", which is of ",len(Array),"Length\n")

	// Aadha aadhura array
	aadhura := [3] string{"Satyam"}
	fmt.Print("This is an aadhura array ",aadhura )
	fmt.Println(", which is of ",len(aadhura),"Length\n")

	// pura array
	full := [4] int{11,22,33,44}
	fmt.Print("This is an array ",full )
	fmt.Println(" which is of ",len(full),"Length\n")

	// special index pai element only
	special := [100] int{ 7:7, 77:77}
	fmt.Print("This is an special array ",special )
	fmt.Println(" which is of ",len(special),"Length\n")

	// normal for
	for i:=0 ; i<len(special) ; i++{
		if(special[i]!= 0){
			fmt.Println("hey",special[i],"is at",i,"index in array but its position is",i+1)
		}
	}

	// for -range loop
	i := 0 
	for index,value := range special{
		if(value == 0){
			continue
		}
			fmt.Println("hey",value,"is at",index,"index in array but its position is",index+1)
		
		i++
	}


}