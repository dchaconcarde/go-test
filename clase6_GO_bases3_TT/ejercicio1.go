package main
import (
	"fmt"
)

func main() {

	usuario := new(Usuario)
	usuario.Nombre = "Daniel"
	usuario.Apellido = "Chacon"
	usuario.Edad = 23
	usuario.Correo = "danielchacon@gmail.com"
	usuario.Contrasena = "lasuperpasssecreta"

	fmt.Printf("Nombre: %s \n",usuario.Nombre)
	fmt.Printf("Apellido: %s \n", usuario.Apellido)
	fmt.Printf("Edad: %d\n", usuario.Edad)
	fmt.Printf("Correo: %s\n", usuario.Correo)
	fmt.Printf("Contraseña: %s\n", usuario.Contrasena)

	fmt.Printf("\n \n \n")

	usuario.cambiarNombre("Juan", "Perez")
	usuario.cambiarEdad(43)
	usuario.cambiarCorreo("juanperez@gmail.com")
	usuario.cambiarContrasena("laMegasupercontrasena")

	fmt.Printf("\n \n \n")

	fmt.Printf("Nombre: %s \n",usuario.Nombre)
	fmt.Printf("Apellido: %s \n", usuario.Apellido)
	fmt.Printf("Edad: %d\n", usuario.Edad)
	fmt.Printf("Correo: %s\n", usuario.Correo)
	fmt.Printf("Contraseña: %s\n", usuario.Contrasena)

}

type Usuario struct {
	Nombre string
	Apellido string
	Edad int
	Correo string
	Contrasena string
}

func (u *Usuario) cambiarNombre(nombre, apellido string){
	u.Nombre = nombre;
	u.Apellido = apellido
	fmt.Printf("Su nombre y apellidos han sido cambiados a: %s %s\n", u.Nombre, u.Apellido)
}

func (u *Usuario) cambiarEdad(edad int){
	u.Edad = edad
	fmt.Printf("Su edad ha sido cambiada a %d\n", u.Edad)
}

func (u *Usuario) cambiarCorreo(correo string){
	u.Correo = correo
	fmt.Printf("Su correo ha sido cambiado a %s\n", u.Correo)

}

func (u *Usuario) cambiarContrasena(contrasena string){
	u.Contrasena = contrasena
	fmt.Printf("Contraseña cambiada :)\n")
}
