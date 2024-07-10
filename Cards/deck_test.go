package main

import (
	// "io/ioutil"
	"os"
	"testing"
)

func TestNewDeck( t *testing.T){
	deck:= newDeck()
	if len(deck)!=24{
		t.Errorf("Expect 24 cards in deck, but got %v",len(deck))
	}
	for i:=0 ; i<len(deck) ;i++{
		for j:=i+1 ; j<len(deck) ; j++{
			if deck[i]==deck[j]{
				t.Errorf("Same card occure : %v",deck[i])
			}
		}
	}
}

func TestWritingAndReadingDeck(t *testing.T){
	os.Remove("_testDeck")

	deck := newDeck()

	deck.saceToFile("_testDeck")

	cards := readCardsFile("_testDeck")

	if len(cards)!=24 {
		t.Errorf("Expected 24 cards in deck, but got  : %v",len(cards))
	} 

	os.Remove("_testDeck")
}