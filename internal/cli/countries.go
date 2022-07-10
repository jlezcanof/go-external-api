package cli

import (
	"fmt"

	countriescli "github.com/jlezcanof/go-external-api/internal"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlagName = "name"
const idFlagIsFullText = "idFlagIsFullText"

// InitCountriesCmd initialize beers command
func InitCountriesCmd(service countriescli.CountryService) *cobra.Command {
	initCountriesCmd := &cobra.Command{
		Use:   "countries",
		Short: "Print data about api rest countries",
		Run:   runCountriesFn(service),
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
		} else {
			fmt.Println("no ha introducido flag name")
		}

		fmt.Println("fulltext is", fullText)
		service.GetCountriesPrint()

	}

}
