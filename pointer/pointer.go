package main

import (
	"fmt"
)

func main() {

	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *int32

	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println("puntero pointerX", pointerX)
	fmt.Println("puntero pointerY", pointerY)
	fmt.Println("puntero pointerZ", pointerZ)
	fmt.Println("puntero &pointerZ", &pointerZ)
	pointerZ = &x
	fmt.Println("puntero pointerZ", pointerZ)
	fmt.Println("puntero &pointerZ", &pointerZ)
	fmt.Println("puntero *pointerZ", *pointerZ)

	fmt.Println("----------")
	type Animal struct {
		name  int
		edad  int
		color int
	}
	var perro = new(Animal)
	fmt.Println(perro == nil)
	fmt.Println(*perro)

	fmt.Println("----------")

	type Casa struct {
		valor       int
		ubicacion   int
		coordenadas int
	}

	miCasa := &Casa{
		valor: 1, coordenadas: 3,
	}

	miCasa.valor = 3
	fmt.Println("*miCasa", *miCasa)
	fmt.Println("&miCasa", &miCasa)
	fmt.Println("miCasa", miCasa)

	miCasa1 := &Casa{
		valor: 7, coordenadas: 3,
	}
	fmt.Println("*miCasa1", *miCasa1)
	fmt.Println("&miCasa1", &miCasa1)
	fmt.Println("----------")

}
