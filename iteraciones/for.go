package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 100; i += 5 {
		if i == 50 {
			continue
		}

		if i == 60 {
			break
		}

		fmt.Println(i)
	}
}
