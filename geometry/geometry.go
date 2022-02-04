package geometry

import (
	"log"
)

var boundaryConditions string
var dim int
var L []int
var V int
var NN [][]int

func InitGeometry(d int, ell []int, bc string) {
	if d < 1 || d > 4 {
		panic("Dimensionality has to be bewtween 1 and 4 (included)")
	}
	boundaryConditions = bc
	dim = d
	L = ell
	V = 1
	for _, value := range L {
		V *= value
	}
	NN = make([][]int, 2*dim)
	for nu := 0; nu < 2*dim; nu++ {
		NN[nu] = make([]int, V)
	}
	switch dim {
	case 1:
		initD1()
	case 2:
		initD2()
	case 3:
		initD3()
	case 4:
		initD4()
	}
	log.Println("Sono in InitGeometry e il volume è ", V)
	log.Println("Sono in InitGeometry e l'array di L è ", L)
}

func initD1() {
}
func initD2() {
}
func initD3() {
}
func initD4() {
}
