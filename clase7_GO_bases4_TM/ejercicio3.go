package main
import(
	"fmt"
)

func main() {
	salary := 30000
	if salary < 150000{
		fmt.Println(fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de %d", salary))
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}	
}