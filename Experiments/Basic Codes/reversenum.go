package main

func revnum(i int) int{
	newnum := 0
	for i!=0{
		last := i%10
		newnum = newnum*10 + last
		i = i/10 
	}
	return newnum
}