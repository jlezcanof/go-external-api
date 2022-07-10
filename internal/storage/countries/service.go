package countries

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
	response, err := http.Get(fmt.Sprint("%v%v", c.url, countriesEndpoint))
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &countries)
	if err != nil {
		return nil, err
	}

	return
}

// other impl
func (c *countryService) GetCountriesPrint() (err error) {
	fmt.Println("invoke get countries print")
	response, err := http.Get(fmt.Sprintf("%v%v", c.url, countriesEndpoint))
	if err != nil {
		return err
	}
	fmt.Println(response)

	return
}
