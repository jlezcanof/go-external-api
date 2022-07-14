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
	//countryEndpoint   = "/v3.1/name/%s?fullText=%t"
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
	//resp, _ := http.Get(fmt.Sprintf("%v%v", c.url, countriesEndpoint))
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body)) //de esta manera si que saca las keys
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

	//var result countriescli.Country
	err = json.Unmarshal(contents, &countries)
	if err != nil {
		return nil, err
	}

	return
}

func (c *countryService) GetOneCountry(name string, isFulltext bool) (country countriescli.Country, err error) {
	fmt.Println("invoke get one country")

	url := fmt.Sprintf("%v/v3.1/name/%s?fullText=%t", c.url, name, isFulltext)
	fmt.Println(url)

	responseCountry, error := http.Get(url)

	if error != nil {
		//return nil, nil
		fmt.Printf(fmt.Sprintf("hubo error en la invocacion get: %v", error))
	}

	contents, err := ioutil.ReadAll(responseCountry.Body)
	if error != nil {
		fmt.Println("hubo error en la lectura de la response")
		//return nil, error
	}

	fmt.Println(&contents)
	error = json.Unmarshal(contents, &country)
	if error != nil {
		fmt.Println("hubo error al un marshallear el objeto json de la response")
		//return nil, error
	}

	return
}
