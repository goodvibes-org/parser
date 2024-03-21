package ingredientes

import (
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/xuri/excelize/v2"
)

func return_remap() map[string]string {
	return map[string]string{
		"COSING Ref No":                 "COSING Ref No",
		"INCI name":                     "INCI name",
		"INN name":                      "INN name",
		"Ph. Eur. Name":                 "Ph. Eur. Name",
		"CAS No":                        "CAS No",
		"EC No":                         "EC No",
		"NamexCas":                      "NamexCas",
		"NamexEC":                       "NamexEC",
		"NamexName":                     "NamexName",
		"Mix":                           "Mix",
		"Anexo.iii.name":                "Anexo.iii.name",
		"Anexo.iii.EC":                  "Anexo.iii.EC",
		"Anexo.iii.CAS":                 "Anexo.iii.CAS",
		"Chem/IUPAC Name / Description": "Chem/IUPAC Name / Description",
		"nchar":                         "nchar",
		"Synonyms.formatx":              "synonyms",
		"-":                             "-",
		"Name to Compare \"Tool\" (Risk databases)": "actual_name",
		"Restriction":                 "Restriction",
		"Function":                    "Function",
		"Anexo.iii.Criteria":          "Anexo.iii.Criteria",
		"Info para Reporte":           "info_para_reporte",
		"Update Date":                 "Update Date",
		"Observaciones":               "Observaciones",
		"Citas":                       "cita",
		"Group.Cancer":                "Group.Cancer",
		"Ref.Cancer":                  "Ref.Cancer",
		"Volume.Cancer":               "Volume.Cancer",
		"Year.Cancer":                 "Year.Cancer",
		"Add Info.Cancer":             "Add Info.Cancer",
		"Add Info.Dev":                "Add Info.Dev",
		"Ref.Dev":                     "Ref.Dev",
		"Group.Endoc":                 "Group.Endoc",
		"Ref.Toxicity.Allergies":      "Ref.Toxicity.Allergies",
		"Add Info.Toxicity.Allergies": "Add Info.Toxicity.Allergies",
		"Add Info.Total/partial use restrictions": "Add Info.Total/partial use restrictions",
		"Ref.Endoc":                          "Ref.Endoc",
		"Ref.Total/partial use restrictions": "Ref.Total/partial use restrictions",
		"Ref.Env":                            "Ref.Env",
		"Add Info.Env":                       "Add Info.Env",
		"Cancer.Risk":                        "cancer_risk",
		"Development.Risk":                   "development_risk",
		"Allergies.Risk":                     "allergies_risk",
		"Endocryne.Risk":                     "endocryne_risk",
		"Prohibited.Risk":                    "prohibited_risk",
		"Env.Risk":                           "env_risk",
		"Total.Risk":                         "total_risk",
	}
}

func ParseIngredientes(fileName string, sheetName string) (bool, error) {

	remap := return_remap()
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		return false, err
	}
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return false, err
	}
	df := dataframe.LoadRecords(rows, dataframe.HasHeader(true))
	for oldName, newName := range remap {
		df = df.Rename(newName, oldName)
	}
	file_ingredientes, err := os.Create("bpc_ingredientes_proc.csv")
	if err != nil {
		return false, err
	}
	df.WriteCSV(file_ingredientes)

	return true, nil

}
