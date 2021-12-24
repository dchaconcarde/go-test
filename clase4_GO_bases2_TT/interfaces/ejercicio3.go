package main

import (
	"fmt"
)

type tiendaModa struct {
	nombre string
	precio float64
	tienda string
	url    string
}
type tiendacolumbia struct {
	tienda string
	precio float64
}
type Ecommerce interface {
	Precio(size string) float64
	Envio(direccion string) string
}

func (t tiendaModa) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return t.precio
	case "mediano":
		var porcentaje float64
		porcentaje = (t.precio / 100) * 3
		return t.precio + porcentaje
	case "grande":
		var porcentaje float64
		flete := 2500
		porcentaje = (t.precio / 100) * 6
		return t.precio + porcentaje + float64(flete)
	}
	return t.precio
}

func (t tiendaModa) Envio(dir string) string {
	enviado := "Enviando un paquete a " + dir
	return enviado
}
func (t tiendacolumbia) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return t.precio
	case "mediano":
		var porcentaje float64
		porcentaje = (t.precio / 100) * 3
		return t.precio + porcentaje
	case "grande":
		var porcentaje float64
		flete := 2500
		porcentaje = (t.precio / 100) * 6
		return t.precio + porcentaje + float64(flete)
	}
	return t.precio
}
func (t tiendacolumbia) Envio(dir string) string {
	enviado := "Enviando un paquete a " + dir
	return enviado
}
func nuevaTienda(mailType string) Ecommerce {
	if mailType == "tiendamoda" {
		return tiendaModa{nombre: "mesa", precio: 2000, tienda: "sucursal Colombia", url: "http//:tiendaModa.com"}
	}
	if mailType == "tiendacolumbia" {
		return tiendaModa{tienda: "sucursal Mexico", precio: 5000}
	}
	return nil
}
func main() {
	fmt.Println("------------------------------------------------")
	tiendaModa := nuevaTienda("tiendamoda")
	precioModa := tiendaModa.Precio("grande")
	envioModa := tiendaModa.Envio("Avenida Independencia 578")
	fmt.Println(tiendaModa)
	fmt.Println(envioModa)
	fmt.Println("Precio del producto:", precioModa)
	fmt.Println("------------------------------------------------")
	tiendaColumbia := nuevaTienda("tiendacolumbia")
	precioColumbia := tiendaColumbia.Precio("Mediano")
	envioColumbia := tiendaColumbia.Envio("Avenida Independencia 578")
	fmt.Println(tiendaColumbia)
	fmt.Println(envioColumbia)
	fmt.Println("Precio del pruducto:", precioColumbia)
	fmt.Println("------------------------------------------------")
}
