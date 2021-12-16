package main

import "fmt"



func main() {
	nombre:="daniel"
	fmt.Println("Cantidad de letras", len(nombre))

	for _, element := range nombre{
		fmt.Println(string(element))
	}
}