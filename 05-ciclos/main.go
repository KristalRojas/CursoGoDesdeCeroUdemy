package main

import "fmt"

func main() {
	//En go solo existe For

	//For mas simple
	// i := 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//For mas comun(?)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	//For forever con un break
	// numero := 0
	// for {
	// 	fmt.Println("Continuo")
	// 	fmt.Println("Ingrese numero secreto: ")
	// 	fmt.Scanln(&numero)
	// 	if numero == 100 {
	// 		break
	// 	}
	// }

	//For con Continue
	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n Valor de i: %d", i)
	// 	if i == 5 {
	// 		fmt.Printf(" multiplicamos por 2 \n")
	// 		i = i * 2
	// 		continue
	// 	}
	// 	fmt.Printf("                 Paso por aqui \n")
	// 	i++
	// }

	//Output de for CON CONTINUE
	// Valor de i: 0                 Paso por aqui
	// Valor de i: 1                 Paso por aqui
	// Valor de i: 2                 Paso por aqui
	// Valor de i: 3                 Paso por aqui
	// Valor de i: 4                 Paso por aqui
	// Valor de i: 5 multiplicamos por 2

	//For SIN Continue
	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n Valor de i: %d", i)
	// 	if i == 5 {
	// 		fmt.Printf(" multiplicamos por 2 \n")
	// 		i = i * 2
	// 	}
	// 	fmt.Printf("                 Paso por aqui \n")
	// 	i++
	// }

	//Output For SIN CONTINUE
	// Valor de i: 0                 Paso por aqui
	// Valor de i: 1                 Paso por aqui
	// Valor de i: 2                 Paso por aqui
	// Valor de i: 3                 Paso por aqui
	// Valor de i: 4                 Paso por aqui
	// Valor de i: 5 multiplicamos por 2
	// 				Paso por aqui

	//La diferencia esta en que si no esta el continue, sale del if y mostra el mensaje que esta en el for
	//en cambio, el que tiene continue, vuelve al inicio del for(?)

	//For con GOTO (goto)
	//Se asigna un nombre a un procedimiento(?), luego con goto se redirige a ese punto para correr de nuevo ese codigo
	var i int = 0

RUTINA:
	println("i vale: ", i)
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a RUTINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
