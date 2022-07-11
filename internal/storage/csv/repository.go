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

var headers = []string{"Region", "SubRegion"}

func (c *csvRepo) SaveDocument(countries *[]countriescli.Country, filename string) error {

	csvFile, err := os.Create(filename)

	if err != nil {
		csvFile.Close()
		return err
	}

	csvWriter := csv.NewWriter(csvFile)

	if err = csvWriter.Write(headers); err != nil {
		csvWriter.Flush()
		csvFile.Close()
		return err
	}

	for _, country := range *countries {
		if err = csvWriter.Write([]string{country.Region, country.SubRegion}); err != nil {
			csvWriter.Flush()
			csvFile.Close()
			return err
		}
	}

	csvWriter.Flush()
	csvFile.Close()

	return nil
}
