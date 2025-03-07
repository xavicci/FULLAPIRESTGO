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
	fmt.Println("----------")
	fmt.Println("----------")

	type Config struct {
		Port        string
		JWTSecret   string
		DatabaseUrl string
	}

	type Server interface {
		Config() *Config
	}

	type Broker struct {
		config *Config
		router string
	}

	mibroker := &Broker{
		config: &Config{
			Port:        "3131",
			JWTSecret:   "sds",
			DatabaseUrl: "CARAMEL",
		},
		router: "AB",
	}

	configuracion := mibroker.config.DatabaseUrl

	otrobroker := &Broker{}
	otrobroker.config.JWTSecret = "SECRET JWT"

	fmt.Println("*mibroker", *mibroker)
	fmt.Println("&mibroker", &mibroker)
	fmt.Println("configuracion", configuracion)
	fmt.Println(" otrobroker.config.JWTSecret")
	fmt.Println("----------")
	fmt.Println("----------")
	fmt.Println("----------")
}
