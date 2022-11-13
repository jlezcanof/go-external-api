package api

import (
	"testing"
)

// func BenchmarkGetCountries(b *testing.B) {
// 	service := NewCountryService()
// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		service.GetCountries("https://restcountries.com/v3.1/all")
// 	}
// }

func TestGetCountries(t *testing.T) {

	service := NewCountryService()
	var url = "https://restcountries.com/v3.1/all"

	_, err := service.GetCountries(url)
	if err != nil {
		t.Fatalf("got an error %v", err)
	}
}
