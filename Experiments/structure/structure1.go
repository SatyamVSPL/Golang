package main

import "fmt"


type contactInfo struct{
	address string
	mobile int
}

type persons struct{
	firstName string
	lastName string
	contactInfo
}


func main() {
	
	satyam :=persons{ 
		firstName: "Satyam",
		lastName: "Yendhe",
		contactInfo: contactInfo{
			address: "xyz",
			mobile: 1234567890,
		},

	}

	satyam.print()

	satyam.updateFirstName("jaduu")
	satyam.print()

	updateLastName(&satyam , "Boo")
	satyam.print()


}

func (details persons) print(){
	fmt.Printf("%+v\n",details)
}

func (details *persons) updateFirstName(newName string){
	details.firstName = newName
}

func updateLastName(details *persons, newName string){
	details.lastName = newName
}