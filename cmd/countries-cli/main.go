package main

import (
	"fmt"

	countriescli "github.com/jlezcanof/go-external-api/internal"
	"github.com/jlezcanof/go-external-api/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("probando")

	var service countriescli.CountryService

	rootCmd := &cobra.Command{Use: "countries-cli"}
	//rootCmd.AddCommand(cli.InitCountriesCmd())
	rootCmd.AddCommand(cli.InitCountriesCmd(service))
	rootCmd.Execute()

	fmt.Println("finalizando")
	fmt.Println("finalizando 2")
}
