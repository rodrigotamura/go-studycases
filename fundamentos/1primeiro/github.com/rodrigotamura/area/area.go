package area // coloca o nome do pacote

import "math" // Pi é uma proporção definida pela relação entre o perímetro
// de uma circunferência e seu diâmetro
const Pi = 3.1416 // constante pública - 1a letra maiuscula

// Circ é resp. por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a área de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

// Náo visível para fora!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
