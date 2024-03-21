package productos

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/xuri/excelize/v2"
)

func ParseProductos(filename string) {
	fileHandle, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := fileHandle.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	df := dataframe.LoadRecords(rows, dataframe.HasHeader(true))
	var validID = regexp.MustCompile(`ingredient\s\d*`)
	var ingredintesNames []int
	var restoNames []int
	names := df.Names()
	for idx, str := range names {
		matched := validID.MatchString(str)
		if matched {
			ingredintesNames = append(ingredintesNames, idx)
		} else {
			restoNames = append(restoNames, idx)
		}
		if str == "rubro" {
			df = df.Rename("rubro_id", str)
		} else {
			df = df.Rename(strings.ToLower(str), str)
		}
	}
	first_column := df.Col("descripcion")
	solo_ingredientes_cols := df.Select(ingredintesNames)
	ingredientes := solo_ingredientes_cols.Mutate(first_column)
	resto := df.Select(restoNames)
	file_ingredientes, err := os.Create("bpc_productos_proc_ingredientes.csv")
	if err != nil {
		fmt.Println("Error creando el archivo de salida de ingredientes")
		return
	}
	file_productos, err := os.Create("bpc_productos_proc.csv")
	if err != nil {
		fmt.Println("Error creando el archivo de salida de productos")
	}
	ingredientes.WriteCSV(file_ingredientes)
	resto.WriteCSV(file_productos)
}
