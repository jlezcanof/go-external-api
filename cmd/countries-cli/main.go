package main

import (
	countriescli "github.com/jlezcanof/go-external-api/internal"
	"github.com/jlezcanof/go-external-api/internal/cli"
	"github.com/jlezcanof/go-external-api/internal/storage/api"
	"github.com/jlezcanof/go-external-api/internal/storage/csv"

	"os"

	"github.com/spf13/cobra"

	//"net/http/pprof"
	"runtime/pprof"
)

func main() {

	var apiService countriescli.CountryService
	var csvService countriescli.CSVService

	apiService = api.NewCountryService()
	csvService = csv.NewCSVService()

	rootCmd := &cobra.Command{Use: "country-cli"}
	rootCmd.AddCommand(cli.InitCountriesCmd(apiService, csvService))
	rootCmd.Execute()

	// memory profiling code starts here
	f, _ := os.Create("countries.mem.prof")
	defer f.Close()
	pprof.WriteHeapProfile(f)
}
