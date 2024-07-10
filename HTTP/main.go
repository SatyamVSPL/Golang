package main

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
	"os"
	
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil{
		fmt.Println("Error :",err)
		os.Exit(1)
	}


	// bs := make([]byte , 99999)

	// req.Body.Read(bs)
	// fmt.Println(string(bs))

	wC := writerCustom{}
	io.Copy(wC,req.Body)

	
}

// custom writer function
type writerCustom struct{}

func (writerCustom) Write(b []byte) (int ,error){
	fmt.Println(string(b))
	fmt.Println("Total no. of bytes are :",len(b))
	return len(b),nil
}