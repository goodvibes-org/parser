package main

import (
	"fmt"

	"github.com/goodvibes-org/parser/ingredientes"
)

func main() {
	fmt.Println("Todavia nada mucho por aqui")
	// productos.ParseProductos()
	retVal := ingredientes.ParseIngredientes("holaa.xlsx")
	fmt.Println(retVal)
}
