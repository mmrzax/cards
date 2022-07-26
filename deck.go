package cards

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck Create a new type of Deck
// Which is a slice of strings
type Deck []string // Deck is like kind of slice or a subclass or extend in OO languages

func NewDeck() Deck {
	cards := Deck{}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join(d, ",")
}

func (d Deck) SaveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1: Log the error and return a call to newDeck()
		// Option #2: Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return ss
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
