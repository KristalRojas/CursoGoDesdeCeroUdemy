package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Kristal Rojas")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
	//Al momento de realizar la ultima accion del main (en este caso era ingresar algo por teclado),
	//automaticamente terminara la func miNombreLento, el runtime no queda esperando a que finalice,
	//por eso las goroutines se usan con los channels
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
