package main

import (
	"fmt"
	"io"
	"os"
)

type logData struct{
	file *os.File
}

func (lD *logData) Read(data []byte) (int , error){
	// fmt.Println("Number of Bytes : ",len(data))
	return lD.file.Read(data)
}

func main() {
	filename:=os.Args[1]

	data  ,err := os.Open(filename) 
	if err!=nil {
		fmt.Println("This is error",err)
		os.Exit(1)
	}
	// fmt.Printf("%T",data)
	// io.Copy(os.Stdout, data)
	lD := logData{file:data}
	io.Copy(os.Stdout, &lD)

}
