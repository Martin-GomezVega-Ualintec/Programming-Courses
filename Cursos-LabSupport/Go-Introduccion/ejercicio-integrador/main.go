package main

import (
	"fmt"

	"github.com/Martin-GomezVega-Ualintec/Programming-Courses/Cursos-LabSupport/Go-Introduccion/ejercicio-integrador/models"
)

// Desarrollar un programa que recorra un conjunto de Puntos(x e y), por cada elemento debe verificar
// si x es mayor a y, en ese caso se debe hacer swap de los valores utilizando una función que reciba x e y,
// realice el intercambio y no retorne nada, por último imprimir cada elemento.

// A tener en cuenta:
// 		La struct de punto debe ser importada en el main
//		Se debe usar al menos una función
// 		Se deben usar punteros

func main() {
	//puntos := []int{1, 2, 3, 1, 2, 4, 5}

	puntos := models.Puntos{
		X: 2,
		Y: 1,
	}

	for _, j := range puntos {
		(&puntos).CompararValores(j.x, j.y)
	}

	//puntos.print()
}

func (pointerToPuntos *Puntos) CompararValores() {
	if (*&pointerToPuntos).x > (*&pointerToPuntos).y {
		aux := *&pointerToPuntos.x
		(*&pointerToPuntos).x = (*&pointerToPuntos).y
		(*&pointerToPuntos).y = aux
	}
}

func (pointerToPuntos *Puntos) CompararValores2(x int, y int) {
	if x > y {
		aux := x
		x = y
		y = aux
	}
}

func (p Puntos) print() {
	fmt.Printf("%+v", p)
}

func main2() {
	queue := make(chan struct {
		string
		int
	})
	go sendPair(queue)
	pair := <-queue
	fmt.Println(pair.string, pair.int)
}

func sendPair(queue chan struct {
	string
	int
}) {
	queue <- struct {
		string
		int
	}{"http:...", 3}
}

// https://stackoverflow.com/questions/13670818/pair-tuple-data-type-in-go
