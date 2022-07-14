package countriescli

// Countries representation of api result into data struct
type Countries struct {
	Countries *[]Country `json:"result"`
}

// Country representation of country into data struct
type Country struct {
	Region    string `json:"region"`
	SubRegion string `json:"subregion"`
	Name      Name   `json:"name"`
	Fifa      string `json:"fifa"`
}

type Name struct {
	Common     string     `json:"common"`
	Oficial    string     `json:"official"`
	NativeName NativeName `json:"nativeName"`
}

type NativeName struct {
	Spa Spa `json:"spa"`
}

type Spa struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

// CountryService definition of methods to access a data country (api rest)
type CountryService interface {
	GetCountries() ([]Country, error)
	GetOneCountry(name string, isfullText bool) (Country, error)
}

type CSVService interface {
	SaveDocument(*[]Country, string) error
}

// TODO deleteNewCountry initialize struct country
/*func NewCountry(region string, subregion string) (c Country) {
	c = Country{
		Region:    region,
		SubRegion: subregion,
	}
	return
}*/
