package main

import "fmt"

type phone string

func (p phone) check() bool{
	number := string(p)
	if len(number)==10{
		// fmt.Println(len(number))
		return true
	}else{
		return false
	}
}

type info struct{
	address string
	mobile phone
	zipcode int
}

type employees struct{
	 name string
	 age int
	 information info
}

func main() {
	satyam := employees{
		name:"Satyam", 
		age:22, 
		information: info{
			address: "Tahen",
			mobile:  "1234567890",
			zipcode: 123456,
		},
	}

	if satyam.information.mobile.check(){
		fmt.Println("Mobile no. is valid")
	}else{
		fmt.Println("Mobile no. is not valid")
	}

	fmt.Println(satyam.name," your age is",satyam.age)

	satyam.age = 23
	fmt.Println("New age :",satyam.age)

	fmt.Println(satyam)
}