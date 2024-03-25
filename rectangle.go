package main

// Rectangle representa un rectángulo con ancho y alto
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calcula el área del rectángulo
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calcula el perímetro del rectángulo
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// IsSquare verifica si el rectángulo es un cuadrado
func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

func main() {
	// Código principal
}
