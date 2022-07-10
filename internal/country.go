package countriescli

// Countries representation of api result into data struct
type Countries struct {
	Countries *[]Country `json:"result"`
}

// Country representation of country into data struct
type Country struct {
	Region    string `json:"region"`
	SubRegion string `json:"subregion"`
	//Name string `json:"name"`
	//ProductID int       `json:"product_id"`
	//Price     string    `json:"price"`
	//BeerID    int       `json:"beer_id"`
	//Category  string    `json:"category"`
	//Type      *BeerType `json:"type"`
	//Brewer    string    `json:"brewer"`
	//Country   string    `json:"country"`
}

// CountryService definition of methods to access a data country (api rest)
type CountryService interface {
	GetCountries() ([]Country, error)
	GetOneCountry() (Country, error, name string, isFulltext bool)
}

type CSVService interface {
	SaveCSV(*[]Country, string) error
}

// NewCountry initialize struct country
func NewCountry(region string, subregion string) (c Country) {
	c = Country{
		Region:    region,
		SubRegion: subregion,
	}
	return
}
