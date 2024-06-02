package funciones

import "fmt"

func Exponential(value int) {
	if value > 1000000000 {
		return
	}
	fmt.Println(value)
	Exponential(value * 2)
}
