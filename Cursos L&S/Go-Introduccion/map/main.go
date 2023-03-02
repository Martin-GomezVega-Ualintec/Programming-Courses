package main

import "fmt"

func main() {
	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#4bf745",
		}
		fmt.Println(colors)
	*/

	colors := make(map[string]string)
	colors["white"] = "#fff"
	colors["green"] = "#4bf745"
	//delete(colors, "white")
	printMap((colors))
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, " is ", hex)
	}
}
