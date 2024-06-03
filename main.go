package main

import (
	"fmt"

	"github.com/salatielosorno/Go/goroutines"
)

/* import custom package */

func main() {
	/*
		estado, texto := variables.ConviertoATexto(1588)
		fmt.Println(estado)
		fmt.Println(texto)
	*/

	/*
		os := runtime.GOOS

		if os == "Linux." || os == "OS X." {
			println("Esto no es Mac, es ", os)
		} else {
			println("Esto es Mac ;)")
		}
		switch os {
		case "Linux.":
			fmt.Println("Esto el Linux")
		case "darwin":
			fmt.Println("Esto es darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/

	/*
		numberInt, message := ejercicios.Evaluate("100")
		fmt.Println("Number:", numberInt)
		fmt.Println(message)
	*/
	/*
		teclado.IngresoNumeros()
	*/
	/*
		iteraciones.Iterar()
	*/

	/*
		fmt.Println(ejercicios.MultiplicationTable())
	*/
	/*
		files.GrabaTabla()
	*/

	/*
		files.SumaTabla()
	*/
	/*
		files.LeoArchivo()
	*/
	/*
		funciones.LlamarClosure()
	*/

	/*
		funciones.Exponential(2)
	*/

	/*
		arreglos_slices.MuestroArreglos()
	*/
	/*
		arreglos_slices.MuestroSlice()
	*/
	/*
		arreglos_slices.Capacidad()
	*/

	/*
		mapas.MostrasMapas()
	*/

	/*
		users.AltaUsuario()
	*/

	/*
		pedro := new(modelos.Hombre)
		ejercicios.HumanosRespirando(pedro)
		martha := new(modelos.Mujer)
		ejercicios.HumanosRespirando(martha)
	*/

	/*
		defer_panic.VemosDefer()
	*/

	/*
		defer_panic.EjemploPanic()
	*/

	go goroutines.MiNombreLentooo("Salatiel")

	fmt.Println("Estoy aquí")

	/*
		Esta línea hace que se espere un
		valor por teclado. De esta forma el código
		asíncrono puede ejecutarse.
	*/
	var x string
	fmt.Scanln(&x)
}
