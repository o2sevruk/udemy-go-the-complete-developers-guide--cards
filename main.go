package main

import "fmt"

func main() {
	cards := []string{"Five of Spades", newCard()}
	cards = append(cards, "Two of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of spades"
}
