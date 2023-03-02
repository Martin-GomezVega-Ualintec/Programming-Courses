package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Crear un nuevo tipo de 'deck'
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// 0666 -> cualquiera puede leer y escribir el archivo
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// retorna byteSlice y una variable de error (es nulo si no tiene error)
	bytS, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		// Nos indica que salió algo mal mientras corriamos el programa al asignarle un 1
		os.Exit(1)
	}
	s := strings.Split(string(bytS), ",")
	return deck(s)
}

// Devolver al azar y que no se repita cada vez que ejecutamos
func (d deck) shuffle() {
	// para iniciar a una hora distinta time.Now().UnixNano()
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
