package main

import "fmt"

func main() {
	//Una manera de crear un mapa es con make
	// paises := make(map[string]string)
	paises := make(map[string]string, 5) //<- Agregar int al final, se le esta diciendo la cantidad de elementtos
	//asi reserva espacio, no significa que este llegue solo hasta ahi, si no que usa menos recursos
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	//Otra forma de
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	//Agregar items a map
	campeonato["River Plate"] = 25
	fmt.Println(campeonato)

	//Update item in a maps
	campeonato["Chivas"] = 55
	fmt.Println(campeonato)

	//Borrar item en un map
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	//Recorremos el mapa
	//Equipo y puntaje seria como key value, dado que el range siempre devuelve ambos items
	//Range por alguna razon siempre devuelve los datos desordenados, no suelen estar ordenados siempre igual
	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	//Ahora buscamos el valor de un item en el map, a apartir de su key, en este caso la key seria los
	//nombres de los equipos de futbol, se retorna dos valores, el valor que se busca y un booleano, para
	//saber si existe o no el item
	// puntaje, existe := campeonato["Mineiro"]
	//Se imprime, %t : el verbo t es para los booleanos <-------
	// fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
	//OUTPUT: El puntaje capturado es 0, y el equipo existe false
	//En este caso se estaba buscando al Mineiro, que no se encuentra en el map, por lo cual muestra
	//los zero values de puntaje int y existe bool

	//Ahora buscaremos uno que exista en el map
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
	//OUTPUT: El puntaje capturado es 55, y el equipo existe true
	//Ahora si muestra valores

}
