package main

import "fmt"

//CREACION DE ARRAY GLOBAL
// var tabla [10]int

//Creacion de Matriz
// var matriz [5][7]int

func main() {
	// tabla[0] = 1
	// tabla[5] = 15

	//Tambien se puede crear y asignar valor de inmediato
	tabla := [10]int{5, 6, 98, 1, 4, 0, 3, 65, 8, 99}

	for i := 0; i < len(tabla); i++ {

		fmt.Println(tabla[i])
	}

	//Print de matriz
	// matriz[3][5] = 1
	// fmt.Println(matriz)

	//Slices
	//Un slice se declara cuando no se le asigna la longitud []int (No se pone el argumento del largo)
	matriz := []int{2, 5, 4}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}

func variante2() {
	//Si tiene largo, entonces es un vector
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}

//Creando slices con make
func variante3() {
	//Make - comando especial para crear slices y adicionar elementos en ejecucion
	//Se declara un slice ya que no se asigna valor al la longitud ([]), siguiente numero es el largo con el que
	//que iniciara el slice, y el segundo numero 20, es una asignacion de la capacidad maxima del slice
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

//***OJO**** Printf va concatenando todos los printf del doc, por eso de usa \n para hacer un salto de linea

//Usando Append para agregar datos a slices
func variante4() {
	//Se crea un slice que tiene capacidad 0 y maximo 0
	nums := make([]int, 0, 0)

	//Se agregara 100 veces el numero i, el append se usan cuando ya te quedas sin espacio en el slice
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
	//OUTPUT: Largo 100, Capacidad 128

	//tiene un largo de 100 y una capacidad maximo de 128, esta ultima la asigna go automaticamente al
	//leer el largo, la diferencia entre no asignar largo y capacidad, es que cuando esta señalado, go ya
	//tiene un espacio en memoria esperando por algo, en cambio cuando este super el maximo asignado,
	//go asignara espacio en memoria para esos items que se agregan, por ende baja la performance y ya no
	//es tan rapido como si hubiera estado todo listo para funcionar
	//En funciones pequeñas no se logra notar una diferencia, pero en grandes aplciaciones hay diferencia
}
