package main

import "log"

func main() {
	//El archivo prueba.txt no existe, por lo que daria un error
	// archivo := "prueba.txt"
	// f, err := os.Open(archivo)

	// //defer: indica que esa sentaencia se ejecutara solo al final de la funcion, perfecto para cerrar dbs
	// defer f.Close()

	// if err != nil {
	// 	fmt.Println("error abriendo el archivo")
	// 	os.Exit(1)
	// }

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		//recover: trae el resultado del ultimo panic, en el ejemplo se utiliza defer para cuando pase por
		//un panic, al ultimo se ejecute el defer con el recover dentro. Cuando no hay panic, recover sera nil
		reco := recover()
		if reco != nil {
			//Fatalf: es como printF, pero graba directamente en archivo de log, puede formatear textos
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco) //verbo %v: valor variable (objeto)
			//output log.Fatalf: 2021/07/28 17:39:12 Ocurrio un error que genero Panic
			//Al ser un log, automaticamente toma hora y fecha
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}

}
