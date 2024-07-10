package main

// import "fmt"

func revStr(str string) string {
	// fmt.Println(len(str))
	var rev string = ""
	for i:= len(str)-1 ; i>=0 ;i--{
		rev += string(str[i])
	}
	return rev
}