package ejercicios

import "strconv"

func Evaluate(value string) (int, string) {
	numberToEvaluate, err := strconv.Atoi(value)

	if err != nil {
		return 0, "Ocurrió un error :( \n" + err.Error()
	}

	if numberToEvaluate > 100 {
		return numberToEvaluate, "Es mayor a 100"
	} else {
		return numberToEvaluate, "Es menor a 100"
	}
}
