package main
import "fmt"

func main() {
	var employees = map[string]int{"Benjamin":20, "Nahuel":26, "Brenda":19, "Dario":44, "Pedro":30}
	var moreThan21 int;
	
	for _, age := range employees {
		if age >= 21{
			moreThan21++;
	}

	}
	fmt.Println(employees)
	
	fmt.Println("Empleados con más de 21 años: ", moreThan21)
	
	employees["Federico"] = 25;

	fmt.Println("Se contrató a Federico")

	delete(employees, "Pedro")

	fmt.Println("Pedro ha sido despedido")

	fmt.Println(employees)

}