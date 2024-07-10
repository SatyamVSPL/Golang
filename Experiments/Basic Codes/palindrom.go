package main

func plaindromeNum(i int) string {
	if(i == revnum(i)){
		return "Plaindrome"
	}else{
		return "Not Palindrome"
	}
}

func plaindromeStr(i string) string{
	if(i == revStr(i)){
		return "Plaindrome"
	}else{
		return "Not Palindrome"
	}
}