package main

// import "fmt"


func main() {
	// var card string = "Ace of Spades"
	// card := newDeck()

	// // // Replace loop with custom data type
	// // for _,value := range card{
	// // 	fmt.Println(value)
	// // }

	// print(card)

	// fmt.Println("\n\nThis is first hand")
	// // hand , remaningCards := deal(card, 4)

	// // print(hand)
	// // print(remaningCards)

	// fmt.Println("\nThis is deck converted to string : ")
	// fmt.Println(card.toSring())

	cards := newDeck()
	cards.print()
	// hand,Remain := deal(cards,4)
	// print(hand)
	// print(Remain)

	// hand.saceToFile("MyCards")

	// print(readCardsFile("MyCards"))

	cards.shuffle().print()
}

