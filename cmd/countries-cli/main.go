package main

import (
	"fmt"

	countriescli "github.com/jlezcanof/go-external-api/internal"
	"github.com/jlezcanof/go-external-api/internal/cli"
	countries "github.com/jlezcanof/go-external-api/internal/storage/api"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("probando")

	var apiService countriescli.CountryService
	//var csvService countriescli.CSVService

	apiService = countries.NewCountryService()

	//csvService = countries.newCsvService()

	rootCmd := &cobra.Command{Use: "countries-cli"}
	rootCmd.AddCommand(cli.InitCountriesCmd(apiService))
	rootCmd.Execute()

}
