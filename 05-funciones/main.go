package main

import "fmt"

func main() {
	//En go todo son func, los metodos son funciones
	fmt.Println(uno(5))

	//Se puede inicializar variables desde el return de una funcion
	numero, estado := dos(1)

	fmt.Println(numero)
	fmt.Println(estado)

	//Prueba de una funcion con n cantidad definida de parametros (de un mismo tipo)
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(2, 23, 45, 68))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 46, 17, 25, 26, 98, 17))
}

//func uno, retorna un dato
// func uno(numero int) int {
// 	return numero * 2
// }

//func uno, retorna un dato, de una forma distinta
//Se encapsula el valor de retorno junto con el nombre de la variable, luego no es necesario poner
//la variable en el retorno, ya que previamente esta señalado cual es la variable a retornar
func uno(numero int) (z int) {
	z = numero * 2
	return
}

//func dos retorna dos datos de distinto tipo, para retornar una lista de datos, se deben encapsular entre parentesis EJ:(int, bool)
func dos(numero int) (int, bool) {
	if numero == 1 {
		return numero, false
	} else {
		return numero * 2, true
	}
}

//En go, hasta el momento no cuenta con funciones con sobrecarga, es decir que no pueden haber dos funciones
//que tengan el mismo nombre pero con distintos parametros.

//Funcion que recibe n cantidad de enteros, al poner ... antes del tipo de dato, estamos declarando
//que esa funcion recibira datos de un tipo, pero no sabe cuantos, en el ej. indica que los parametros
//de entrada son x cualquier numero
func Calculo(numero ...int) int {
	total := 0

	//Instruccion range devuelve siempre dos valores, este toma un rango de parametros que se asigna al lado
	//derecho de range, ese dato debe ser una lista, o un vector, objecto iterable
	//Range devuelve el numero de elemento(index(?)), en este caso no lo necesito, para ignorar tal elemento
	//se utiliza _ en lugar de la variable que no quiero. _ indica que no se usara, y no usa espacio ni nada
	// for _, num := range numero {
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
