package api

import "testing"

func BenchmarkGetCountries(b *testing.B) {
	service := NewCountryService()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.GetCountries()
	}
}

// ojo que hay otro metodo get countries para testear
// func BenchmarkGetOneCountry(b *testing.B) {
func BenchmarkGetOneCountry(b *testing.B) {
	service := NewCountryService()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		service.GetOneCountry("spain", false)
	}
}
