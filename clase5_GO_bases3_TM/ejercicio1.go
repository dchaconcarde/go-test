package main
import (
	"fmt"
	"os"
	"encoding/csv"
)
/*
Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, 
separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

func main() {

	p1 := Productos{1, 200.00, 2}
	p2 := Productos{2, 300.00, 8}
	p3 := Productos{3, 120.00, 12}
	p4 := Productos{4, 80.00, 7}

	productos := []Productos{}

	productos = append(productos, p1, p2, p3, p4)


	csvFile, err := os.Create("products.csv")

	if err != nil {
		fmt.Print("Error: ", err)
	}
	productsTitle:=[]string{}
	productTitle:=fmt.Sprint("ID,\t Precio,\t Cantidad;")
	productsTitle = append(productsTitle,productTitle)
	csvwriter := csv.NewWriter(csvFile)

	_ = csvwriter.Write(productsTitle)

	fmt.Println(productsTitle)
 
	for _, producto := range productos {
		productToPrint := []string{}
		toPrint:=fmt.Sprintf("%d,\t %10.2f,\t %d ;", producto.Id, producto.Precio, producto.Cantidad)
		productToPrint = append(productToPrint, toPrint)
	_ = csvwriter.Write(productToPrint)
		fmt.Println(toPrint)
	}

	csvwriter.Flush()
	csvFile.Close()
}

type Productos struct {
	Id int
	Precio float64
	Cantidad int
}