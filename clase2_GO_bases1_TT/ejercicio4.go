package main
import "fmt"
	

func main() {
	mes:= 1
	switch mes {
		case 1:
			fmt.Println("Enero")
		case 2:
			fmt.Println("Febrero")
		case 3:
			fmt.Println("Marzo")
		case 4:
			fmt.Println("Abril")
		case 5: 
			fmt.Println("Mayo")
		case 6:
			fmt.Println("Junio")
		case 7:
			fmt.Println("Julio")
		case 8:
			fmt.Println("Agosto")
		case 9:
			fmt.Println("Septiembre")
		case 10:
			fmt.Println("Octubre")
		case 11:
			fmt.Println("Noviembre")
		case 12:
			fmt.Println("Diciembre")
		default:
			fmt.Println("Elija un mes correcto")	
	}

	var monthMap = map[int]string{1:"enero", 2:"febrero", 3:"marzo", 4:"abril", 5:"mayo", 6:"junio", 7:"julio", 8:"agosto", 9:"septiembre", 10:"octubre", 11:"noviembre", 12:"diciembre"}
	

	for i, element:= range monthMap{
		if(i == mes){
			fmt.Println(element)
		}
	}

}