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

	servicios = append(servicios, instalacion, reparacion)

	mantenimientoGeneral := Mantenimiento{"MantenimientoGeneral", 200.00}
	superMantenimientoGeneral := Mantenimiento{"SuperMantenimientoGeneral", 800.00}

	mantenimiento = append(mantenimiento, mantenimientoGeneral, superMantenimientoGeneral)

	totalProductos := <-sumarProductos(producto)
	totalServicios := <-sumarServicios(servicios)
	totalMantenimientos := <-sumarMantenimiento(mantenimiento)

	fmt.Println("El total de precio de productos es :", totalProductos)
	fmt.Println("El total de precio de servicios es :", totalServicios)
	fmt.Println("El total de precio de mantenimiento es :", totalMantenimientos)

	fmt.Println("Finalizo el programa")

}

func sumarProductos(productos []Productos) <-chan float64 {
	channel1 := make(chan float64)
	total := 0.0
	go func() {
		for _, p := range productos {
			totalPorProducto := p.Precio * float64(p.Cantidad)
			total += totalPorProducto
		}
		channel1 <- total
	}()
	return channel1
}

func sumarServicios(servicios []Servicios) <-chan float64 {
	channel2 := make(chan float64)
	total := 0.0
	go func() {
		for _, s := range servicios {
			halfHour := float64(s.MinutosTrabajados / 30)
			var rounded int = int(math.Ceil(halfHour))
			total += s.Precio * float64(rounded)
		}
		channel2 <- total
	}()
	return channel2
}

func sumarMantenimiento(mantenimiento []Mantenimiento) <-chan float64 {
	channel3 := make(chan float64)
	total := 0.0
	go func() {
		for _, m := range mantenimiento {
			total += m.Precio
		}
		channel3 <- total
	}()
	return channel3
}

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}
