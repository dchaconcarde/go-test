package main
import "fmt"
import "github.com/skelterjohn/go.matrix"

func main() {

	m1 := Matrix{2, 2, 0, false}
	m1.Altura = 2;
	m1.Anchura = 2;
	setMatrix:=setMatrix(1, 2, 3, 4, m1.Altura, m1.Anchura)
	m1.MaxValue = printMatrix(setMatrix)
	fmt.Println("Altura: ", m1.Altura)
	fmt.Println("Anchura: ", m1.Anchura)
	fmt.Println("Valor maximo: ", m1.MaxValue)
	if m1.Altura == m1.Anchura {
		m1.isCuadratic = true
	}
	fmt.Println("Es una matriz cuadrática: ", m1.isCuadratic)
}

type Matrix struct {
	Altura int
	Anchura int
	MaxValue float64
	isCuadratic bool
}

func setMatrix(valor1, valor2, valor3, valor4 float64, altura, anchura int) *matrix.DenseMatrix {
	a:= matrix.MakeDenseMatrix([] float64{valor1, valor2, valor3, valor4}, altura, anchura)
	return a;
}

func printMatrix(ma *matrix.DenseMatrix) float64 {
	r := ma.Rows()
	c := ma.Cols()
	var maxValue float64 = -9999

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			a := ma.Get(i, j)
			if a>maxValue{
				maxValue = a
			}
			print(" | ",int(a), " ")
		}
		fmt.Println("|")
	}
	return maxValue
}
