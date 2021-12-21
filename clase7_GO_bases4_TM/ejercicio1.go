package main
import(
	"fmt"
)

func main() {
	salary := 30000
	myError := myError{}
	if salary < 150000{
		fmt.Println(myError.Error())
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}	
}

type myError struct {

}

func (e *myError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}