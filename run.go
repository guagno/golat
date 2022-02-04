package main

import (
	"log"
	"github.com/guagno/golat/geometry"
)

func main() {
	geometry.InitGeometry(2, []int{8, 4}, "periodic")
	log.Println("Sono in main: L = ", geometry.L)
}
