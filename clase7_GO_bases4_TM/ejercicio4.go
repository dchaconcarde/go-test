package main
import (
	"fmt"
	"errors"
)

func main() {
	trabajador1 := new(Trabajador)
	trabajador1.Horas = 2000
	trabajador1.PrecioHora = 300
	trabajador1.MesesTrabajados = 6

	salarioTotal, err := salarioPorHorasTrabajadas(trabajador1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El salario total del mes es: %10.2f\n",salarioTotal)

	aguinaldo, err := calcularAguinaldo(salarioTotal, trabajador1.MesesTrabajados)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El aguinaldo para los meses trabajados es: %10.2f\n ", aguinaldo)
}

type Trabajador struct{
	Horas int
	PrecioHora float64
	MesesTrabajados int
}

func salarioPorHorasTrabajadas(trabajador *Trabajador) (float64, error){

	salarioTotal := 0.00
	var precioParaImpuestos float64 = 150000

	if trabajador.Horas < 80 {
		return 0, errors.New("Error: el trabajador no puede haber trabajado menos de 80 horas")
	}
	salarioTotal = float64(trabajador.Horas) * trabajador.PrecioHora
	if salarioTotal >= precioParaImpuestos {
		impuestos := salarioTotal*10/100
		salarioTotal -= impuestos
		fmt.Printf("Se retiró un 10%%= %10.2f, tu salario total quedó en: %10.2f\n", impuestos, salarioTotal)
	}
	return salarioTotal, nil
}

func calcularAguinaldo(salarioDelMes float64, mesesTrabajados int)(float64, error){
	if mesesTrabajados>6{
		mesesTrabajados = 6
		fmt.Println("Trabajó más de 6 meses, se tomaran sólo 6 meses para calcular el aguinaldo")
	}
	if salarioDelMes<0 || mesesTrabajados<0{
		return 0, fmt.Errorf("Error: Los números ingresados no pueden ser negativos, Salario :%10.2f, Meses trabajados: %d", salarioDelMes, mesesTrabajados )
	}
	aguinaldo := ((salarioDelMes/12)*float64(mesesTrabajados))
	return aguinaldo, nil
}
	


