package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goodvibes-org/parser/ingredientes"
	"github.com/goodvibes-org/parser/productos"
	"github.com/urfave/cli/v2"
)

func main() {
	var productos_filename string
	var ingredientes_filename string
	var productos_sheetname string
	var ingredintes_sheetname string

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "productos",
				Aliases: []string{"p"},
				Usage: "Procesa xlsx de productos para generar un archivo csv adecuado para el resto del  software" +
					". Por defecto utiliza `BPC_Productos.xlsx` como archivo de entrada y `Productos` como nombre de  sheet",
				Action: func(ctx *cli.Context) error {
					ok, err := productos.ParseProductos(productos_filename, productos_sheetname)
					if ok {
						fmt.Printf("Se convirtió correctamente la sheet %v del archivo %v", productos_sheetname, productos_filename)
					}
					if err != nil {
						errString := fmt.Sprintf("Error => %v", err)
						return cli.Exit(errString, 1)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "filename",
						Value:       "BPC_Productos.xlsx",
						Usage:       "Elegir nombre de Xlsx de Productos",
						Destination: &productos_filename,
					},
					&cli.StringFlag{
						Name:        "sheetname",
						Value:       "Productos",
						Usage:       "Elegir nombre de sheet",
						Destination: &productos_sheetname,
					},
				},
			},
			{
				Name:    "ingredientes",
				Aliases: []string{"i"},
				Usage: "Procesa xlsx de productos para generar un archivo csv adecuado para el resto del  software" +
					". Por defecto utiliza `BPC_Ingredientes.xlsx` como archivo de entrada y `Ingredientes_Formatted_V1` como nombre de  sheet",
				Action: func(ctx *cli.Context) error {
					ok, err := ingredientes.ParseIngredientes(ingredientes_filename, ingredintes_sheetname)
					if ok {
						fmt.Printf("Se convirtió correctamente la sheet %v del archivo %v", ingredintes_sheetname, ingredientes_filename)
					}
					if err != nil {
						errString := fmt.Sprintf("Error => %v", err)
						return cli.Exit(errString, 1)
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "filename",
						Value:       "BPC_Ingredientes.xlsx",
						Usage:       "Elegir nombre de XLSX de Ingredientes",
						Destination: &ingredientes_filename,
					},
					&cli.StringFlag{
						Name:        "sheetname",
						Value:       "Ingredientes_Formatted_V1",
						Usage:       "Elegir nombre de sheet",
						Destination: &ingredintes_sheetname,
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
