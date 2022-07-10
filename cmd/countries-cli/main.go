package main

import (
	countriescli "github.com/jlezcanof/go-external-api/internal"
	"github.com/jlezcanof/go-external-api/internal/cli"
	"github.com/jlezcanof/go-external-api/internal/storage/api"
	"github.com/jlezcanof/go-external-api/internal/storage/csv"

	"github.com/spf13/cobra"
)

func main() {

	var apiService countriescli.CountryService
	var csvService countriescli.CSVService

	apiService = api.NewCountryService()
	csvService = csv.NewCSVService()

	rootCmd := &cobra.Command{Use: "country-cli"}
	rootCmd.AddCommand(cli.InitCountriesCmd(apiService, csvService))
	rootCmd.Execute()

}
