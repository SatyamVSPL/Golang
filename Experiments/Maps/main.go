package main

import "fmt"

func main() {
	// Initilizing empty map a 
	// var a = map[string]string{}

	a:= map[string]string{}
	fmt.Printf("Type of DS : %+v\n",a)

	a["Name"] = "Audi"
	a["Color"] = "Black"
	a["Year"] ="2019"

	// Accessing particular element
	fmt.Println("Name of car is :",a["Name"])
	fmt.Println(a)

	// addding new element
	a["Model"] = "Xs"
	fmt.Println(a)

	// Updating Element
	a["Model"] = "X100"
	fmt.Println(a)

	// deleting element
	delete(a,"Model")
	fmt.Println(a)


	// Checking for specific element in map
	value , status := a["Name"]
	if status {
		fmt.Println(value,"is present in map")
	}else{
		fmt.Println("Value is not present in map")
	}

	// Itteratring over Map
	for key, value := range a{
		fmt.Println(key,"is key for",value)
	}


	// Iterarting over map in specific order
	b:=[]string{"Color","Year","Name"}

	fmt.Println("\nFunction k sath specific sequence bhi print kiya hai")
	printMap(a,b)
}

func printMap(m map[string]string , order []string ) {
	for _,value := range order{
		fmt.Printf("%v is key for %v\n",value,m[value])
	}
}