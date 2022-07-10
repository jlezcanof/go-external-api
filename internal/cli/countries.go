package cli

import (
	"fmt"
	"log"

	countriescli "github.com/jlezcanof/go-external-api/internal"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlagName = "name"
const idFlagIsFullText = "idFlagIsFullText"

// InitCountriesCmd initialize beers command
func InitCountriesCmd(apiService countriescli.CountryService) *cobra.Command {
	initCountriesCmd := &cobra.Command{
		Use:   "countries",
		Short: "Print data about api rest countries",
		Run:   runCountriesFn(apiService),
	}

	initCountriesCmd.Flags().StringP(idFlagName, "n", "", "name of the country")
	initCountriesCmd.Flags().BoolP(idFlagIsFullText, "f", false, "flag is the search is full text")

	return initCountriesCmd
}

func runCountriesFn(service countriescli.CountryService) CobraFn {

	return func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(idFlagName)
		fullText, _ := cmd.Flags().GetBool(idFlagIsFullText)

		if name != "" {
			fmt.Println("name is", name)
			//country, _ := service.GetOneCountry(name, fullText)
		} else {
			fmt.Println("no ha introducido flag name")
		}
		fmt.Println("fulltext is", fullText)

		countries, err := service.GetCountries()
		if err != nil {
			log.Fatalf("Error while retrieving all countries: %s", err)
		}
		fmt.Println(countries)
		/*
			// de aqui en adelante sacarlo en un metodo de una interface
			csvFile, errorFile := os.Create("output.csv")

			if errorFile != nil {
				csvFile.Close()
				return errorFile
			}

			//fmt.Println("run countries fn ", countries) //solo esta sacando el value, no el key
		*/
	}

}
