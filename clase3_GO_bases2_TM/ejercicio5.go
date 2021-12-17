package main
import "fmt"
import "errors"

func main() {
	const (
		perro = "perro"
		gato = "gato"
		hamster = "hamster"
		tarantula = "tarantula"
	)

	animalPerro, err := animal(perro)
	animalGato, err := animal(gato)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)

	var cantidadTotal float64

	cantidadTotal += animalPerro(5)
	cantidadTotal += animalGato(2)
	cantidadTotal += animalHamster(15)
	cantidadTotal += animalTarantula(40)

	if err != nil {
		fmt.Printf("Error:", err)
	}

	fmt.Println("Cantidad total de comida a comprar:", cantidadTotal, "Kgs")


}

func animal(animal string) (func(float64) float64, error) {
	switch animal {
		case "perro": 
			return func(unidades float64) float64 {
				var result float64=unidades * 10
				return result
			}, nil
		case "gato":
			return func(unidades float64) float64 {
				var result float64=unidades * 5
				return result
			}, nil
		case "hamster":
			return func(unidades float64) float64 {
				var result float64=unidades * 0.25
				return result
			}, nil
		case "tarantula":
			return func(unidades float64) float64 {
				var result float64=unidades * 0.15
				return result
			}, nil
		default:
			return nil, errors.New("Este animal no está en la lista")
	}
}