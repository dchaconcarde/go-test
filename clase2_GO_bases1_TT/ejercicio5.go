package main
import "fmt"

func main() {
	names := []string{"Benjamin", "Nahuel", "Brenda",
	 "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	 fmt.Println(names)

	 names = append(names, "Gabriela")

	 fmt.Println(names)

}