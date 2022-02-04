package main

import (
	"log"
	"github.com/guagno/golat/geometry" geo
)

func main() {
	geometry.InitGeometry(2, {4, 4})
	log.Println("Sono in main: L = ", geometry.L)
}
