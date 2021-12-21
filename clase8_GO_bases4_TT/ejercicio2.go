package main
import (
	"fmt"
	"os"
	"math/rand"
	"errors"
)

func main() {
	client1 := Client{iDGenerator(), "Daniel", "Chacon", 97989832, "Av. Estanislao Lopez"}
	
	client2 := Client{iDGenerator(), "Juan", "Perez", 0, "Av. Italia"}

	leerTxt()

	client1isblank, err := isBlank(client1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	client2isblank, err := isBlank(client2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	clients :=[]Client{}

	if !client1isblank{
		clients = append(clients, client1)
	}

	if !client2isblank{
		clients = append(clients, client2)
	}

	for _, client := range clients {
		fmt.Println(client)
	}



	

}	

type Client struct {
	Legajo int
	Nombre string
	Apellido string
	NumeroTelefono int
	Domicilio string
}

func iDGenerator() int{
	legajo := rand.Intn(100)
	if(legajo == 0){
		panic("Abortando programa, un int nill wtf")
	}
	return legajo
	
}

func leerTxt(){
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	file, err := os.Open("customers.txt")
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println(file)
}

func isBlank(client Client) (bool, error){
	if(client.Legajo == 0 || client.NumeroTelefono == 0){
		return true, errors.New("No puede tener valores en 0")
	}else if(client.Nombre == "" || client.Apellido == "" || client.Domicilio == ""){
		return true, errors.New("No puede dejar campos vacíos")
	}
	return false, nil
}