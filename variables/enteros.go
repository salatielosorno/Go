package variables

import "fmt"

func MuestroEnteros() {
	/* Declarativa */
	var intcomun int /* value is 0, in Go every var start with the minimum value of the type */
	/* Por asignación*/
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
