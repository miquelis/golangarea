package golangarea

import "math"

//Pi é uma proporção númerica
const Pi = 3.1416

// Circ é responsável por calcular a área da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é reponsável por calcular a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
