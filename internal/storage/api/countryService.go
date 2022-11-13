package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	countriescli "github.com/jlezcanof/go-external-api/internal"
	jsoniter "github.com/json-iterator/go"
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

/* new*/
func (c *countryService) GetCountries(url string) (countries []countriescli.Country, err error) {
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
		fmt.Println("hubo error al unmarshallear el objeto json de la response")
		return nil, error
	}

	return
}

/* */

func (c *countryService) standarUnmarshall(data []byte, countries any) error { //countries []countriescli.Country
	err := json.Unmarshal(data, &countries)
	if err != nil {
		return err //TODO crear estructura de errores
		//errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return nil
}

func (c *countryService) betterUnmarshall(data []byte, countries any) error { //countries []countriescli.Country
	err := jsoniter.Unmarshal(data, &countries)
	if err != nil {
		return err //TODO crear estructura de errores
		//errors.WrapDataUnreacheable(err, "can't parsing response into beers")
	}
	return nil
}
