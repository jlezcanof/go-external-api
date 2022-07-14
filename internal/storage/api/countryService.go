package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	countriescli "github.com/jlezcanof/go-external-api/internal"
)

const (
	countriesEndpoint = "/v3.1/all"
	countriesURL      = "https://restcountries.com"
)

type countryService struct {
	url string
}

//NewCountryService fetch beers from api rest
func NewCountryService() countriescli.CountryService {
	return &countryService{url: countriesURL}
}

func (c *countryService) GetCountries() (countries []countriescli.Country, err error) {
	fmt.Println("invoke get all countries")
	response, err := http.Get(fmt.Sprintf("%v%v", c.url, countriesEndpoint))
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//var result countriescli.Countryoracle/database:18.4.0-xe
	err = json.Unmarshal(contents, &countries)
	if err != nil {
		return nil, err
	}

	return
}

func (c *countryService) GetOneCountry(name string, isFulltext bool) (countries []countriescli.Country, err error) {
	fmt.Println("invoke get one country")

	url := fmt.Sprintf("%v/v3.1/name/%s?fullText=%t", c.url, name, isFulltext)

	responseCountries, error := http.Get(url)

	if error != nil {
		fmt.Printf(fmt.Sprintf("hubo error en la invocacion get: %v", error))
		return nil, error
	}

	contents, err := ioutil.ReadAll(responseCountries.Body)
	if error != nil {
		fmt.Println("hubo error en la lectura de la response")
		return nil, error
	}

	error = json.Unmarshal(contents, &countries)
	if error != nil {
		fmt.Println("hubo error al un marshallear el objeto json de la response")
		return nil, error
	}

	return
}
