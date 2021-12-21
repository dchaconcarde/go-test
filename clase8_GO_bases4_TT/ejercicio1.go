package main
import (
	"fmt"
	"os"
)

func main(){
	txtName := "customers.txt"
	leerTxt(txtName)
	fmt.Println("Ejecución finalizada")


}

func leerTxt(name string){
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	file, err := os.Open(name)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println(file)
}
