package main
import "fmt"

func main() {
	edad:=23
	empleado:=true
	antiguedad:=1;
	sueldo:=120000;

	if(edad>=22 && empleado && antiguedad >= 1){
		fmt.Println("Opta por el préstamo")
		if(sueldo>=100000){
			fmt.Println("Prestamo sin interés")
		}else{
			fmt.Println("Prestamo con interés")
		}
	}else{
		fmt.Println("No opta por el préstamo")
	}

}