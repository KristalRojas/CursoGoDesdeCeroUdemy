package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

// Genero Humano
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre //<--Asi se hereda la estructura
	// edad       int
	// altura     float32
	// peso       float32
	// respirando bool
	// pensando   bool
	// comiendo   bool
}

func (h *hombre) respirar() { h.respirando = true }

func (h *hombre) comer() { h.comiendo = true }

func (h *hombre) pensar() { h.pensando = true }

func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Es hombre"
	} else {
		return "Mujer"
	}
}

func (h *hombre) estaVivo() bool { return h.vivo }

func HumanosRespirando(h humano) {
	h.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", h.sexo())
}

// func main() {
// 	Pedro := new(hombre)
// 	Pedro.esHombre = true
// 	HumanosRespirando(Pedro)

// 	Maria := new(mujer)
// 	HumanosRespirando(Maria)
// }

//Genero Animal

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalesRespirar(a animal) {
	a.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoros(a animal) int {
	if a.EsCarnivoro() == true {
		return 1
	}
	return 0
}

// func main() {
// 	totalCarnivoros := 0
// 	Dogo := new(perro)
// 	Dogo.carnivoro = true
// 	AnimalesRespirar(Dogo)
// 	totalCarnivoros = +AnimalesCarnivoros(Dogo)

// 	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)
// }

//Ser vivo
func estoyVivo(s SerVivo) bool {
	return s.estaVivo()
}

func main() {
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)
	fmt.Printf("Estoy vivo %t \n", estoyVivo(Dogo))
}
