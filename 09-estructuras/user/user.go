package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

//Forma de asociar una funcion a una struct
//Me parece que es como un setter
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
