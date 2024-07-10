/*

  *
 * *
* * *

*/

package main

import "fmt"

func main() {
	i := 3
	for row:=0; row<i; row++{
		for space:=row; space<i-space ; space++{
			fmt.Print(" ")
		}
		for star:=0;star<=row ; star++{
			fmt.Print("* ")
		}
		fmt.Println("")
	} 
}