package main
import(
	"fmt"
	"errors"
)

func main() {
	salary := 30000
	if salary < 150000{
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}