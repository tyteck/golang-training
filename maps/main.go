package main

import "fmt"

func main() {

	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	colors["white"] = "#FFFFFF"
	colors["black"] = "#000000"

	colors["foo"] = "foo"
	delete(colors, "foo")

	fmt.Println(colors)

	fmt.Println(colors["red"])

	print(colors)

}

func print(m map[string]string) {
	fmt.Println("--------------------------")
	for colorName, value := range m {
		fmt.Println(colorName, "is", value)
	}
}
