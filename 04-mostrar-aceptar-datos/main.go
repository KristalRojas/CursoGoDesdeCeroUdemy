package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {

	//Como Scannear por consola

	//En windows y mac, Scanf no funciona correctamente, debido a una diferencia entre los enters, y es
	//que el 2 numero no lo lee. Go esta basado en linux, por lo cual ahi si funciona bien
	// fmt.Println("Ingrese numero 1: ")
	// fmt.Scanf("%d",&numero1)

	// fmt.Println("Ingrese numero 2: ")
	// fmt.Scanf("%d",&numero2)

	//Para evitar ese problema, utilziar la funcion Scanln, esta solo necesita el puntero
	// fmt.Println("*Scan*")
	// fmt.Println("Ingrese numero 1: ")
	// fmt.Scanln(&numero1)

	// fmt.Println("Ingrese numero 2: ")
	// fmt.Scanln(&numero2)

	// fmt.Println(numero1 + numero2)

	//Bufio y os
	fmt.Println("*Bufio y Os*")
	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text() //<--- scanner.Text() lee lo que ingresas por teclado y lo guarda en leyenda
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
