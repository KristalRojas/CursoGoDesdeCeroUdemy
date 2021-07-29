package main

import "fmt"

var result int

func main() {
	//Middlewares tienen que ver con seguridad (?)
	fmt.Println("Inicio")
	//llamamos a funcion middleware
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

//La entrada y la salida deben ser iguales, aqui recibe una funcion que toma 2 parametros de entrada y un int de salida
//por lo cual retorna una funcion con los mismos parametros
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
