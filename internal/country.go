package countriescli

//"encoding/json"

// Country representation of country into data struct
type Country struct {
	//Name string `json:"name"`
	Region    string `json:"region"`
	SubRegion string `json:"subregion"`
	//ProductID int       `json:"product_id"`
	//Price     string    `json:"price"`
	//BeerID    int       `json:"beer_id"`
	//Category  string    `json:"category"`
	//Type      *BeerType `json:"type"`
	//Brewer    string    `json:"brewer"`
	//Country   string    `json:"country"`
}

// CountryService definition of methods to access a data country
type CountryService interface {
	GetCountries() ([]Country, error)
	GetCountriesPrint() ([]Country, error)
	GetOneCountry() (Country, error, name string, isFulltext bool)
}

// NewCountry initialize struct country
func NewCountry(region string, subregion string) (c Country) {
	c = Country{
		Region:    region,
		SubRegion: subregion,
	}
	return
}
