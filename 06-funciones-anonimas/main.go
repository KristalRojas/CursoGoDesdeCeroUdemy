package main

import "fmt"

//Creacion de funcion anonima
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo dos numeros (10,5) = %d \n", Calculo(10, 5))

	//La ventaja de esto, es que se puede volver a redefinir la funcion. ej:
	//Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos dos numeros (10,5) = %d \n", Calculo(10, 5))
	//Para redefinir una funcion, hay que tener en cuenta que se deben respetar los paarmetros de entrada y salida

	//Dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos dos numeros (10,5) = %d \n", Calculo(10, 5))

	Operaciones()

	//Closures
	fmt.Println("****CLOSURES****")
	tablaDel := 2 //valor que pide la funcion Tabla
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

//Closure (Proteccion e isolation de codigo)
//Se maneja con los closure la isolacion de codigo
//pueden acceder a variables que han sido creadas por fuera de la funcion (no confundir con var globales y publicas,
//mas bien accede a las variables que estan en la funcion de afuera)
//funcion que devuelve una funcion
func Tabla(valor int) func() int {
	numero := valor
	//secuencia=variable que estara solo presente en Tabla
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
	//La funcion de retorno debe coincidir con la func declarada en la func de afuera
}
