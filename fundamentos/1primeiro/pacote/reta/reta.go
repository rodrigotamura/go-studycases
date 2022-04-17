package main

import "math"

// iniciando com letra maiuscula é PÚBLICO (STRUCT, VARIAVEL, CONSTANTE, ETC.)
// público = visível dentro e fora do pacote

// podemos ter vários arquivos dentro do pacote
// pois tudo isso será compilado em uma estrutura
// mas não podemos criar funçoes com mesmo nome

// iniciando com letra MINUSCULA é privado DENTRO DO PACOTE

// o nome do arquivo não influencia em nada, mas sim o nome do pacote

// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// Algo público necessita de comentário

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return // não declarando nada pois estamos retornando nomeado
}

// Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
