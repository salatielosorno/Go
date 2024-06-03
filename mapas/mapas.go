package mapas

import "fmt"

func MostrasMapas() {
	paises := make(map[string]string)

	// fmt.Println(paises)

	paises["México"] = "CDMX"
	paises["Argentina"] = "Buenos aires"

	/* fmt.Println(paises)
	fmt.Println(paises["Argentina"]) */

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	// Print values from map, they are shown unordered
	/* for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d\n", equipo, puntaje)
	} */

	// Delete element
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	// Find a value
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
