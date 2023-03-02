package main

//import "fmt"

//var card string

func main() {
	//card = "H"

	//var card string = "Ace of Spades"

	//var card string
	//card = "Ace of Spades"

	//card := newCard()
	//card = "c"
	//fmt.Println(card)

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades ")
	cards := newDeck()

	//cards.print()
	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Fice of Diamonds"
}
