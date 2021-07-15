package main

import (
	"fmt"
	"strconv" //<--Paquete para convertir
)

//Como se declaran las variables (globales ya que estan fuera del main)
var numero int
var texto string
var status bool

//Dentro de numericos
// float32 float64
// int8 int16 int32 int64
// uint8 uint16 uint32 uint64

func main() {
	//Creacion de variable
	// var numero2 int
	//Creacion de variable 2 forma
	// := este simbolo crea e instancia una variable, solo se puede usar al crear la variable, no asi para asignar
	//parecido a python, no se le esta pasando el tipo de dato, pero Go lo asume (creo que no siempre(?), pero al menos con string, int y float si)
	// numero3 := 4

	//Crear varias variables en una sola linea
	// var num1, num2, num3 int

	//Crear y asignar valor en una sola linea, 3 variables con el mismo valor
	// num1, num2, num3 := 2

	//Crear y asignar valor en una sola linea, 3 variables con distintos valores
	// num1, num2, num3 := 2, 5, 67

	//Crear y asignar valor en una sola linea, 4 variables con distintos valores y tipo de dato
	num1, num2, num3, text, status := 2, 5, 67, "Hola", true

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(text)
	fmt.Println(status) //Este status es true

	//Print zero values, si imprime una variable que "aun no tiene valor asignado", Go mostrara un valor default
	//para los numeros es 0, string es "" (Empty), bool es false
	// fmt.Println(status) //status fue declarada una variable global, pero no se asigno ningun valor, entonces es false

	mostrarStatus()

	//Go no deja que se convierta de un tipo a otro
	//En este caso num2 es de tipo int (se declaro arriba), y se le esta asignando la variable moneda
	//de tipo float32, osea con decimales, esto no puede ser.
	//Ejemplo
	var moneda float32 = 56.32
	// 	num2 = moneda
	//Para convertir o parsear de un numero a otro usar
	num2 = int(moneda)

	fmt.Println(num2)

	//Tampoco permitira la asignacion de un string a un int o float64, ojo que tambien vale para int8 a int16
	// text = moneda
	//Para parsear de numero a texto, hacer lo siguiente:
	text = fmt.Sprintf("%f", moneda) //<-- %f, seguido del porcentaje, viene una letra que seria el verbo
	//Verbo de conversion: f seria float, d para int, investigar mas!!

	//Ahora si podemos imprimir en texto el valor de la variable moneda
	fmt.Println("Text: ", text)

	//Itoa es una funcion muy utilizada en C, convierte un numero int a string, pero recibe solo int generico
	text = strconv.Itoa(int(moneda)) //<--Se parsea de float64 a int generico para usar Itoa

	fmt.Println("Texto Itoa: ", text)
}

func mostrarStatus() {
	//Valor de la variable Global, zero value de bool es false
	fmt.Println(status)
}
