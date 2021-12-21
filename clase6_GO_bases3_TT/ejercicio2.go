package main
import (
	"fmt"
)
/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar 
productos a los usuarios. Para ello requieren que tanto los usuarios como 
los productos tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.

*/
func main() {
	user := new(Usuario)
	user.Nombre = "Daniel"
	user.Apellido = "Chacon"
	user.Correo = "danielchacon@gmail.com"

	user2 := new(Usuario)
	user2.Nombre = "Juan"
	user2.Apellido = "Perez"
	user2.Correo = "juanpe@gmail.com"

	pollo := nuevoProducto("Pollo", 200.00)
	carne := nuevoProducto("Carne", 400.00)
	arroz := nuevoProducto("Arroz", 60.00)

	agregarProducto(user, pollo, 5)
	agregarProducto(user, carne, 2)
	agregarProducto(user, arroz, 8)
	
	agregarProducto(user2, arroz, 10)

	fmt.Printf("Nombre: %s \n",user.Nombre)
	fmt.Printf("Apellido: %s \n", user.Apellido)
	fmt.Printf("Correo: %s\n", user.Correo)
	fmt.Println("Productos: ", user.Productos, "\n\n")

	fmt.Printf("Nombre: %s \n",user2.Nombre)
	fmt.Printf("Apellido: %s \n", user2.Apellido)
	fmt.Printf("Correo: %s\n", user2.Correo)
	fmt.Println("Productos: ", user2.Productos)
	

}

type Usuario struct {
	Nombre string
	Apellido string
	Correo string
	Productos []Productos
}

type Productos struct {
	Nombre string
	Precio float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) Productos {
	p := Productos{}
	p.Nombre = nombre
	p.Precio = precio

	return p

}

func agregarProducto(user *Usuario, producto Productos, cantidad int){
	producto.Cantidad = cantidad
	user.Productos =append(user.Productos, producto)
}

func borrarProductos(user Usuario){
	user.Productos = nil
}