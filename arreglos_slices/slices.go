package arreglos_slices

import (
	"fmt"
)

var tableS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tableS)
	fmt.Println(arreglo)
	porcion := arreglo[3:]   // Slice created using a vector, from position 3 to end
	porcion2 := arreglo[:5]  // Slice created using a vector, from the initial position to 5
	porcion3 := arreglo[6:7] // Slice created using a vector, from 6 to 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo: %d, Capacidad: %d\n", len(nums), cap(nums))
}
