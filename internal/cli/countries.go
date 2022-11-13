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
	initCountriesCmd.Flags().BoolP(idFlagIsFullText, "f", false, "flag full text of search")

	return initCountriesCmd
}

func runCountriesFn(apiService countriescli.CountryService, csvService countriescli.CSVService) CobraFn {

	return func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString(idFlagName)
		fullText, _ := cmd.Flags().GetBool(idFlagIsFullText)

		invokeCommand(name, apiService, csvService, fullText)

	}

}

func invokeCommand(name string, apiService countriescli.CountryService, csvService countriescli.CSVService, fullText bool) {

	var nameFile = "output-all-countries"
	var url = "https://restcountries.com/v3.1/all"
	if name != "" {
		//url para un pais
		nameFile = "output-country-" + name
		url = fmt.Sprintf("https://restcountries.com/v3.1/name/%s?fullText=%t", name, fullText)
	}

	fmt.Println("url es:", url)
	countries, error := apiService.GetCountries(url)
	if error != nil {
		log.Fatalf("Error while retrieving all countries: %s", error)
	}

	csvService.SaveDocument(&countries, nameFile)
}
