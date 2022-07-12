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
func InitCountriesCmd(apiService countriescli.CountryService, csvService countriescli.CSVService) *cobra.Command {
	initCountriesCmd := &cobra.Command{
		Use:   "countries",
		Short: "Print data about api rest countries",
		Run:   runCountriesFn(apiService, csvService),
	}

	initCountriesCmd.Flags().StringP(idFlagName, "n", "", "name of the country")
	initCountriesCmd.Flags().BoolP(idFlagIsFullText, "f", false, "flag is the search is full text")

	return initCountriesCmd
}

func runCountriesFn(apiService countriescli.CountryService, csvService countriescli.CSVService) CobraFn {

	return func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(idFlagName)
		fullText, _ := cmd.Flags().GetBool(idFlagIsFullText)

		fmt.Println("fulltext is", fullText)
		if name != "" {
			fmt.Println("name is", name)
			country, error := apiService.GetOneCountry(name, fullText)
			if error != nil {
				log.Fatalf("Error while retrieving the country : %s", name)
			}
			fmt.Println(country)
		} else {
			fmt.Println("no ha introducido flag name")
		}

		countries, err := apiService.GetCountries()
		if err != nil {
			log.Fatalf("Error while retrieving all countries: %s", err)
		}

		csvService.SaveDocument(&countries, "output-file"+name)
	}

}
