package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./12-manejo-archivos/archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./12-manejo-archivos/archivo.txt")
	defer archivo.Close()
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
}

func graboArchivo() {
	archivo, err := os.Create("./12-manejo-archivos/newarchivo.txt")
	defer archivo.Close()
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
}

func graboArchivo2() {
	fileName := "./12-manejo-archivos/newarchivo.txt"
	if Append(fileName, "\n Esta es una segunda linea") == false {
		fmt.Println("Error en graboArchivo2")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error (Append)")
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error (Append)")
		return false
	}

	return true
}
