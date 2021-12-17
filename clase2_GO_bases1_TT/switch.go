package main
import "fmt"

func main() {
	i:=45
	switch{
	case i< 10:
		fmt.Println("i es menor a 10")
	case i< 50:
		fmt.Println("i es menor a 50")
		fallthrough
	case i > 60:
		fmt.Println("i es mayor a 60")
	default:
		fmt.Println("entro a default")
	}
	//el fallthrough se pasa por los huevos la condición
}