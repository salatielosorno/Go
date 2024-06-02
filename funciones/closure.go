package funciones

import "fmt"

func table(value int) func() int {
	number := value
	secuence := 0

	return func() int {
		secuence++
		return number * secuence
	}
}

func LlamarClosure() {
	tableOf := 5
	MyTable := table(tableOf)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", tableOf, i, MyTable())
	}
}
