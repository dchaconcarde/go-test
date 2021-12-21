package main
import (
	"fmt"
	"math"
)

func main() {

	producto := []Productos{}
	servicios := []Servicios{}
	mantenimiento := []Mantenimiento{}

	iPhone := Productos{"IPhone", 5000.00, 2}
	samsung := Productos{"Samsung", 3000.00, 6}

	producto = append(producto, iPhone, samsung)

	instalacion := Servicios{"Instalacion", 150.00, 40}
	reparacion := Servicios{"Reparacion", 120.00, 120}

	servicios = append(servicios,instalacion, reparacion)

	mantenimientoGeneral := Mantenimiento{"MantenimientoGeneral", 200.00}
	superMantenimientoGeneral := Mantenimiento{"SuperMantenimientoGeneral", 800.00}

	mantenimiento = append(mantenimiento, mantenimientoGeneral, superMantenimientoGeneral)

	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	var sumaProductos float64
	var sumaServicios float64
	var sumaMantenimiento float64

	go func() <-chan int{
		sumaProductos = sumarProductos(producto)
		return channel1
	}()
	
	go func() <-chan int{
		sumaServicios = sumarServicios(servicios)
		return channel2
	}()

	go func() <-chan int{
		sumaMantenimiento = sumarMantenimiento(mantenimiento)
		return channel2
	}()

	fmt.Println("El total de precio de productos es :", sumaProductos)
	fmt.Println("El total de precio de servicios es :", sumaServicios)
	fmt.Println("El total de precio de mantenimiento es :", sumaMantenimiento)

	variable1 := <- channel1
	variable2 := <- channel2
	variable3 := <- channel3

	fmt.Println("Terminó: ",variable1, variable2, variable3)
	


}

func sumarProductos(productos []Productos) float64{
	total := 0.0
	for _, p := range productos {
		totalPorProducto := p.Precio * float64(p.Cantidad)
		total+=totalPorProducto
	}
	return total
}
	 



func sumarServicios(servicios []Servicios) float64 {
	total := 0.0
	for _, s := range servicios {
		halfHour := float64(s.MinutosTrabajados/30)
		var rounded int = int(math.Ceil(halfHour))
		total += s.Precio * float64(rounded)
	}
	return total
}

func sumarMantenimiento(mantenimiento []Mantenimiento) float64 {
	total := 0.0
	for _, m := range mantenimiento{
		total+=m.Precio
	}
	return total
}

type Productos struct {
	Nombre string
	Precio float64
	Cantidad int
}

type Servicios struct {
	Nombre string
	Precio float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}