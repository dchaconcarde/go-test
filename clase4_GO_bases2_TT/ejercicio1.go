package main
import (
    "encoding/json"
    "fmt"
)

func main() {

	p1:= Estudiantes{}

	p1 = Estudiantes{
		Nombre : "Daniel",
		Apellido : "Chacon",
		DNI : 63205180,
		Fecha : "19/06/1997",
	}

	p1.detalle()

	p2:= Estudiantes{"Cecilia", "Suarez", 50026680, "08/07/1996"}

	p2.detalle()

}

type Estudiantes struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}

func (estudiante Estudiantes) detalle(){
	
    estudianteJSON, _ := json.MarshalIndent(estudiante, "", "  ")
     fmt.Printf(string(estudianteJSON))
}