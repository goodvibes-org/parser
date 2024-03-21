package productos

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/xuri/excelize/v2"
)

func cleanData(rows [][]string) [][]string {
	width := len(rows[0])
	for idx, row := range rows {
		if len(row) < width {
			var newrow []string
			newrow = append(newrow, row...)
			for len(newrow) < width {
				newrow = append(newrow, "")
			}
			rows[idx] = newrow
		}
	}
	return rows
}

func ParseProductos(filename string, sheetname string) (bool, error) {
	fileHandle, err := excelize.OpenFile(filename)
	if err != nil {
		return false, err
	}
	rows, err := fileHandle.GetRows(sheetname)
	if err != nil {
		return false, err
	}
	rows = cleanData(rows)
	df := dataframe.LoadRecords(rows, dataframe.HasHeader(true))
	var validID = regexp.MustCompile(`Ingredient\s\d*`)
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
		return false, err
	}
	file_productos, err := os.Create("bpc_productos_proc.csv")
	if err != nil {
		return false, err
	}
	ingredientes.WriteCSV(file_ingredientes)
	resto.WriteCSV(file_productos)
	return true, nil
}
