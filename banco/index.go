package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Variables globales
var (
	opcion   int
	cantidad float32 = 0.0
	cuenta   float32
)

func main() {
	menu()
}

// Menu principal del sistema
func menu() {

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
		IngresarDinero()
	case 2:
		fmt.Println("two")
	case 3:

		ConsultarDinero()
	default:
		fmt.Println("Número no valido")
	}

}

func IngresarDinero() {

	LimpiarCOnsola()

	fmt.Println("\033[31m######################################\033[0m")
	fmt.Println("\u001B[31m#              INGRESO               #\u001B[0m")
	fmt.Println("\u001B[31m######################################\u001B[0m")
	fmt.Println("                                      ")
	fmt.Print("Ingrese la cantidad a ingresar: ")
	fmt.Scan(&cantidad)

	fmt.Println("\u001B[31mCantidad ingresada exitosamente\u001B[0m")
	cuenta = cantidad
}

func ConsultarDinero() {
	LimpiarCOnsola()
	fmt.Println("\033[31m######################################\033[0m")
	fmt.Println("\u001B[31m#              CONSULTA              #\u001B[0m")
	fmt.Println("\u001B[31m######################################\u001B[0m")
	fmt.Println("                                      ")
	fmt.Println("Tienes ", cuenta, " €")
}

func LimpiarCOnsola() {
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
func colores() {
	fmt.Println("\033[31mTexto en rojo\033[0m")     // Rojo
	fmt.Println("\033[32mTexto en verde\033[0m")    // Verde
	fmt.Println("\033[33mTexto en amarillo\033[0m") // Amarillo
	fmt.Println("\033[34mTexto en azul\033[0m")     // Azul
	fmt.Println("\033[35mTexto en magenta\033[0m")  // Magenta
	fmt.Println("\033[36mTexto en cian\033[0m")     // Cian
	fmt.Println("\033[0mTexto normal")              //
}
*/
