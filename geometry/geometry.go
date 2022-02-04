package geometry

import (
	"log"
)

var dim int
var L []int
var V int
var NN [][]int

func InitGeometry(d int, ell []int) {
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
	log.Println("Sono in InitGeometry e il volume è ", V)
	log.Println("Sono in InitGeometry e l'array di L è ", L)
}
