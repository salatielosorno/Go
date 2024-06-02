package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/salatielosorno/Go/ejercicios"
)

var filename = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.MultiplicationTable()
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear archivo. " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
	fmt.Println("File created.")
}

func Append(filename string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append. " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto + "\n")

	if err != nil {
		fmt.Println("Error durante el WriteString. " + err.Error())
		return false
	}

	arch.Close()

	return true
}

func SumaTabla() {
	var texto string = ejercicios.MultiplicationTable()

	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido.")
	}
}

func LeoArchivo() {
	archivo, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error al leer el archivo. " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}
