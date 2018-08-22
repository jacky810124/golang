package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#FFFFFF"
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"

	updateMap(colors, "red", "red")

	printMap(colors)

	// delete key: white
	// delete(colors, "white")

	// colors := map[string]string{
	// 	"red": "#FF0000",
	// }
}

func printMap(colors map[string]string) {
	for key, value := range colors {
		fmt.Println(key, value)
	}
}

func updateMap(colors map[string]string, key string, value string) {
	colors[key] = value
}
