package main

/* import custom package */
import (
	"fmt"

	"github.com/salatielosorno/Go/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
