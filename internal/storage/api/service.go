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

// other impl
func (c *countryService) GetCountries() (countries []countriescli.Country, err error) {
	resp, _ := http.Get(fmt.Sprintf("%v%v", c.url, countriesEndpoint))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body)) //de esta manera si que saca las keys
	//fmt.Println(resp)

	fmt.Println("invoke get countries print")
	response, err := http.Get(fmt.Sprintf("%v%v", c.url, countriesEndpoint))
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

func (c *countryService) GetOneCountry() (Country, error, name string, isFulltext bool) {
	fmt.Println("invoke get one country")

	return
}
