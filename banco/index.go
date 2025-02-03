package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Variables globales
var (
	opcion        int
	cantidad      int    = 0
	cuenta        int    = 0
	controlSalida bool   = false
	nombre        string = "/data/bank.txt"
)

func main() {

	/*
		Verifica si hay archivo y abre uno
	*/
	if ExistText(nombre) == false {
		f, err := os.Create(nombre)
		check(err)

	}

	/*	El ciclo for simula ser un ciclo while
		la variable controlSalida inicia siendo false y cuando el
		usuario decide salir cambia a true
	*/
	for controlSalida == false {
		menu()
		LimpiarCOnsola()
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
Creación del archivo para guardar datos
*/
func ExistText(nombre String) bool {
	f, err := os.Open(nombre)
	if err != nil {
		return false
	}
	f.Close()
	return true
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
	fmt.Println("IV.  Salir      ")
	fmt.Print("Ingrese su opción: ")
	//log.Println("Hora de ingreso")
	//fmt.Scan() lee por consola lo que se escribe por teclado
	fmt.Scan(&opcion)
	switch opcion {
	case 1:
		IngresarDinero()
	case 2:
		RetirarDinero()
		time.Sleep(4 * time.Second)
	case 3:
		ConsultarDinero()
	case 4:
		controlSalida = true
	default:
		fmt.Println("Tecla no válida")
		fmt.Println("Ingrese de nuevo la opción")
		time.Sleep(4 * time.Second)
	}

}

/*
Función que guarda o acumula la cantidad ingresada por el usuario
*/
func IngresarDinero() {
	LimpiarCOnsola()
	fmt.Println("\033[31m######################################\033[0m")
	fmt.Println("\u001B[31m#              INGRESO               #\u001B[0m")
	fmt.Println("\u001B[31m######################################\u001B[0m")
	fmt.Println("                                      ")
	fmt.Print("Ingrese la cantidad a ingresar: ")
	fmt.Scan(&cantidad)
	fmt.Println("\u001B[31mCantidad ingresada exitosamente\u001B[0m")
	cuenta = cuenta + cantidad
	f.writeFile(cuenta)
	defer f.Close()
}

func RetirarDinero() {
	LimpiarCOnsola()
	fmt.Println("\033[31m######################################\033[0m")
	fmt.Println("\u001B[31m#              Retirar               #\u001B[0m")
	fmt.Println("\u001B[31m######################################\u001B[0m")
	fmt.Println("                                      ")
	fmt.Print("Ingrese la cantidad a retirar: ")
	fmt.Scan(&cantidad)

	if cantidad <= 0 || cantidad > cuenta {
		fmt.Println("No puede retirar la cantidad indicada")
		fmt.Println("Saliendo...")
		time.Sleep(2 * time.Second)
	} else {
		cuenta = cuenta - cantidad
		fmt.Println("Cantidad retirada con éxito")
		fmt.Println("Saliendo...")
		time.Sleep(2 * time.Second)
	}

}
func ConsultarDinero() {
	LimpiarCOnsola()

	/*	Usamos la libreria de bufio para poder leer en consola

	 */
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\033[31m######################################\033[0m")
	fmt.Println("\u001B[31m#              CONSULTA              #\u001B[0m")
	fmt.Println("\u001B[31m######################################\u001B[0m")
	fmt.Println("                                      ")
	fmt.Println("Tienes ", cuenta, " €")
	fmt.Println("Presiona cualquier tecla para continuar")

	scanner.Scan()
	input := scanner.Text()

	if input == "p" {
		fmt.Println("Regresando al menu principal......")
		time.Sleep(2 * time.Second)
	}

}

/*
Esta función es muy útil para tener la consola limpia
tiene dos opciones dependiendo del sistema operativo
*/
func LimpiarCOnsola() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		var cmd *exec.Cmd
		cmd = exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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
