package main

import "fmt"
import "errors"

func main() {
	const (
		minimo = "minimo"
		promedio = "promedio"
		maximo = "maximo"
	)

	minimun, err := operacion(minimo)
	promedium, err := operacion(promedio)
	maximun, err := operacion(maximo)

	valorMinimo := minimun(2,3,3,3,4,54,7,4,1,10)
	valorPromedio := promedium(2,3,3,3,4,54,7,4,1,10)
	valorMaximo := maximun(2,3,3,3,4,54,7,4,1,10)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Valor mínimo es:", valorMinimo)
	fmt.Println("Valor promedio es:", valorPromedio)
	fmt.Println("Valor maximo es:", valorMaximo)
}

func minFunc(numeros ...int) float64{
	min:=99999
	for _, numero := range numeros{
		if numero <= min{
			min = numero
		}
	}
	return float64(min)
}

func maxFunc(numeros ...int) float64{
	max:=-99999
	for _, numero := range numeros{
		if numero > max{
			max = numero
		}
	}
	return float64(max)
}

func promFunc(numeros ...int) float64{
	i:=0;
	global:=0
	for _, numero := range numeros{
		global+=numero
		i++
	}
	prom:=float64(global/i)
	return prom

}

func operacion(operacion string) (func(...int) float64, error){
	switch operacion {
		case "minimo":
			return minFunc, nil 
		case "promedio":
			return promFunc, nil
		case "maximo":
			return maxFunc, nil
		default:
			return nil, errors.New("invalid operacion")
	}
}