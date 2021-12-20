package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strings"
)

func main() {

	csvFile, err := os.Open("products.csv")
	if err != nil{
		fmt.Println("Error", err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil{
		fmt.Println("Error", err)
	}
	fmt.Println("ID\t     Precio\t Cantidad")
	for i, line := range csvLines{
	  if i<1{
		  continue
		}
		product := productData{
			ProductoFromCsv : line[0],
		}
		i++
		result := strings.Replace(product.ProductoFromCsv, ",", "", -1)
		result = strings.Trim(result, ";")
		fmt.Println(result)
	}
}

type productData struct {
	ProductoFromCsv string
}