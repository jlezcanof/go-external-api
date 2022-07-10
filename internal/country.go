package countriescli

//"encoding/json"

// Country representation of country into data struct
type Country struct {
	Name string `json:"name"`
	//ProductID int       `json:"product_id"`
	//Price     string    `json:"price"`
	//BeerID    int       `json:"beer_id"`
	//Category  string    `json:"category"`
	//Type      *BeerType `json:"type"`
	//Brewer    string    `json:"brewer"`
	//Country   string    `json:"country"`
}

// CountryService definiton of methods to access a data country
type CountryService interface {
	GetCountries() ([]Country, error)
	GetCountriesPrint() error
}

/* // NewBeer initialize struct beer
func NewBeer(productID int, name, category, brewer, country, price string, beerType *BeerType) (b Beer) {
	b = Beer{
		ProductID: productID,
		Name:      name,
		Category:  category,
		Type:      beerType,
		Brewer:    brewer,
		Country:   country,
		Price:     price,
	}
	return
}*/
