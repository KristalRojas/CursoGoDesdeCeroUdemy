package main

import (
	"fmt"
	"time"

	us "./user"
	//Los paquetes que creemos deben estar por debajo de la carpeta que contiene main.go, en este caso
	//la carpeta user debe estar por debajo de la 09-estructuras por motivos de compilacion
	//Se importa agregando la ruta de la carpeta, no la ruta del .go
	//Importante mantener el nombre de la carpeta al nombre del archivo .go, ej: carpeta user, dentro debe contener user.go
	//La ruta que se importa es de una carpeta ya que eso es un paquete, por eso la ruta es asi
)

//Practicamos herencia dentro de Go
type pepe struct {
	us.Usuario
}

func main() {
	//***********Ignorar 1era parte del cap***************
	//Ejemplo de struct dentro de main
	// User := new(usuario)
	// User.Id = 10
	// User.Nombre = "Pablo"
	// fmt.Println(User)
	//***********Ignorar 1era parte del cap***************

	//En este ejemplo pepe es un puntero de user, osea es una struct que hereda de user
	user := new(pepe)

	user.AltaUsuario(1, "Pablo Tilotta", time.Now(), true)

	fmt.Println(user.Usuario)
}

//***********Ignorar 1era parte del cap***************
//Creacion de struc simple- Solo ejemplo basic dentro de main
//En la realidad no se realizan structs en el main
// type usuario struct {
// 	Id        int
// 	Nombre    string
// 	FechaAlta time.Time
// 	Status    bool
// }
//***********Ignorar 1era parte del cap***************
