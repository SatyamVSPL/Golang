package main

import (
	"fmt"
	// "internal/coverage/rtcov"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := [4]string{"Spade", "Club", "Diamonds", "Hearts"}
	cradValue := []string{"Ace", "One", "Two", "Three", "Four", "Five"}

	for _, suites := range cardSuites {
		for _, values := range cradValue {
			cards = append(cards, suites+" of "+values)
		}
	}
	return cards
}

// func (cards deck) print(){
// 	for index,c_name := range cards{
// 		fmt.Println(index ,c_name)
// 	}
// }

func deal(cards deck, handside int) (deck, deck) {
	return cards[:handside], cards[handside:]
}

func (cards deck) print() {
	for index, c_name := range cards {
		fmt.Println(index, c_name)
	}
}

func (cards deck) toSring() string {

	str := []string(cards)
	return strings.Join(str, ",")
}

func (cards deck) saceToFile(filename string) error {
	return os.WriteFile(filename, []byte(cards.toSring()), fs.FileMode(0666))
}

func readCardsFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	splitStr := strings.Split(string(bs), ",")
	return deck(splitStr)
}

func (d deck) shuffle() deck {
	for i := 0; i < len(d); i++ {

		source := rand.NewSource(time.Now().UnixNano())
		// fmt.Println(source)
		r := rand.New(source)
		randNumber := r.Intn(len(d) - 1)
		// fmt.Println("This is random number :",randNumber)
		// randNumber = int(randNumber)
		d[i], d[randNumber] = d[randNumber], d[i]
		fmt.Println(d[i])

	}
	return d
}
