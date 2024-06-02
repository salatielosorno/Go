package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MultiplicationTable() string {
	scanner := bufio.NewScanner(os.Stdin)

	var number int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese un número:")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid value. Please, try again.")
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d\n", number, i, number*i)
	}

	return texto
}
