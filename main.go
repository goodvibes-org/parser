package main

import (
	"fmt"

	"github.com/goodvibes-org/parser/ingredientes"
	"github.com/goodvibes-org/parser/productos"
)

func main() {
	fmt.Println("Todavia nada mucho por aqui")
	ok, err := productos.ParseProductos("BPC_Productos.xlsx", "Productos")
	oki, erri := ingredientes.ParseIngredientes("BPC_Ingredientes.xlsx", "Ingredientes_Formatted_V1")
	if !ok || !oki {
		fmt.Printf("Hubo un error, si fue ingredientes es :\n%v y si es producto es: \n%v", erri, err)
	} else {
		fmt.Println("EXITO")
	}
}
