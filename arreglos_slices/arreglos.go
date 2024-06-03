package arreglos_slices

import "fmt"

var table [10]int = [10]int{10, 0, 59}

func MuestroArreglos() {
	table[7] = 33
	table[2] = 54

	fmt.Println(table)

	table2 := [10]string{"Pablo", "Juan"}
	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Printf("table[%d]: %d\n", i, table[i])
	}
}
