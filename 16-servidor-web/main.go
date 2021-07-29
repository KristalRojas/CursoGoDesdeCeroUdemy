package main

import "net/http"

func main() {
	//definimos la ruta y la func que debe abrir al ingresasr a dicha ruta, en este caso es index.html
	http.HandleFunc("/", home)
	//Asi se declara el puerto, basico
	http.ListenAndServe(":3000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./16-servidor-web/index.html")
}
