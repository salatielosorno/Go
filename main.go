package main

/* import custom package */
import (
	"fmt"
	"runtime"
)

func main() {
	/*estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/

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
}
