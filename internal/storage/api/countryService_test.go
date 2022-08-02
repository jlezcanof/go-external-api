package api

import "testing"

func BenchmarkGetCountries(b *testing.B) {
	service := NewCountryService()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.GetCountries("https://restcountries.com/v3.1/all")
	}
}
