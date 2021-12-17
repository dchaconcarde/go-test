package main
import "fmt"
import "errors"

func main() {

	prom, err := promedio(10, 5, 12, 10, 5, 8)

	if(err != nil){
		fmt.Println("Ocurrió un error:", err )
	}else{
		fmt.Println("El promedio es:", prom)
	}
}

func promedio(notas ...int) (float64, error){
	i:=0;
	global:=0
	
	for _, nota := range notas {
	  global += nota
	  i++
	}
	prom:=float64(global/i)
	if prom<0 {
		return prom,errors.New("Promedio Negativo")
	}
	return prom, nil

}