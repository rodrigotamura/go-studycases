package matematica

import (
	"fmt"
	"strconv"
)

func Media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	media := total / float64(len(numeros))

	// strconv.ParseFloat -> converter string em float
	mediaArredondada, _ := strconv.ParseFloat(
		fmt.Sprintf("%.2f", media),
		64, // passando que é float64
	)

	return mediaArredondada
}
