package main

import "fmt"

type number []int

func (n number) check(){
	for _,value :=range n{
		if value%2 ==0{
			fmt.Printf("%v is even\n",value)
		}else{
			fmt.Printf("%v is odd\n",value)
		}
	}
}

func main() {
	var n = []int{}
	var n_slice number= n
	n_slice.check()
	
}