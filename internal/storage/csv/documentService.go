package csv

import (
	"encoding/csv"
	"os"

	countriescli "github.com/jlezcanof/go-external-api/internal"
)

type csvRepo struct {
	separator rune
}

func NewCSVService() countriescli.CSVService {
	return &csvRepo{','}
}

var headers = []string{"Region", "SubRegion", "Fifa", "Common country name", "Oficcial country name", "Native name Official", "Native name Common"}

func (c *csvRepo) SaveDocument(countries *[]countriescli.Country, filename string) error {

	// creacion del fichero
	csvFile, err := os.Create(filename)

	// ver si ha habido problemas con la creacion del fichero
	if err != nil {
		csvFile.Close()
		return err
	}

	// preparacion del fichero para su escritura
	csvWriter := csv.NewWriter(csvFile)

	// preparacion del fichero con su cabecera
	if err = csvWriter.Write(headers); err != nil {
		csvWriter.Flush()
		csvFile.Close()
		return err
	}

	// recorremos los pais/es para guardarlo
	for _, country := range *countries {
		if err = csvWriter.Write([]string{country.Region, country.SubRegion, country.Fifa,
			country.Name.Common, country.Name.Oficial, country.Name.NativeName.Spa.Official,
			country.Name.NativeName.Spa.Common}); err != nil {
			csvWriter.Flush()
			csvFile.Close()
			return err
		}
	}

	csvWriter.Flush()
	csvFile.Close()

	return nil
}
