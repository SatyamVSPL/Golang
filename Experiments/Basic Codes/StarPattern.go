/* 

  *
 * *
* * *

*/

package main
// import "fmt"
	
func starpattren(i int) string{
	// var i int
	var Pattren string
	// fmt.Scanln(&i)
	for row:=0 ; row<i ; row++{
		for space:=row ; space<i-1 ; space++{
			Pattren +=" "
		} 
		for star:=0 ; star<=row ; star++{
			Pattren+="* "
		}
		Pattren+="\n"
	}
	return Pattren
}