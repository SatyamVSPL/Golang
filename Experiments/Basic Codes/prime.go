package main

func prime(num int) string {
	for i:=2 ; i<num ; i++{
		if(num%i ==0){
			return "Not Prime"
		}else{
			return "Prime"
		}
	}
	return "idk"
}