package main
import "fmt"


func main() {
	sueldo:=50000
	disc1:=17
	disc2:=10
	fmt.Println("Tu impuesto es:", calcularSalario(sueldo, disc1, disc2))

}

func calcularSalario(sueldo, disc1, disc2 int) float64 {
	var impuesto float64
	if sueldo >= 50000{
		impuesto = float64((sueldo*disc1)/100)
	}else if sueldo >=150000{
		impuesto = float64((sueldo*(disc1+disc2))/100)
	}else{
		impuesto = 0.0
	}
	return impuesto
}