package api

import (
	"errors"
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

	tests := map[string]struct {
		url string
		err error
	}{
		"all the countries":    {url: "https://restcountries.com/v3.1/all", err: nil},
		"obtain country spain": {url: "https://restcountries.com/v3.1/name/spain?fullText=false", err: nil},
		"not exists country":   {url: "https://restcountries.com/v3.1/name/porta?fullText=false", err: errors.New("error")},
	}

	service := NewCountryService()

	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {
			_, err := service.GetCountries(tc.url) //b, err
			if err != nil && tc.err == nil {
				t.Fatalf("not expected any errors and got nil %v", err)
			}

			if err == nil && tc.err != nil {
				t.Error("expected and arror got nil")
			}

			//if tc.want != nil

		})
	}

}
