package main

import (
	"fmt"
)

func main() {
	menu()
}

// Menu principal del sistema
func menu() {

	//variables
	var opcion int
	fmt.Println("######################################")
	fmt.Println("#                BANCO               #")
	fmt.Println("######################################")
	fmt.Println("                                      ")
	fmt.Println("I.   Ingresar dinero ")
	fmt.Println("II.  Retirar dinero ")
	fmt.Println("III. Consultar      ")
	fmt.Print("Ingrese su opción: ")
	//log.Println("Hora de ingreso")

	//fmt.Scan() lee por consola lo que se escribe por teclado
	fmt.Scan(&opcion)
	switch opcion {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Número no valido")
	}

}
