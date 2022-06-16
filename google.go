package gcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindPlaceWithGoogle find coordinate using Google Geolocation API.
func FindPlaceWithGoogle(apiKey, keyword string) ([]Coordinate, error) {
	req, err := http.NewRequest(http.MethodGet, "https://maps.googleapis.com/maps/api/place/findplacefromtext/json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to construct request of Google Geolocation API: %w", err)
	}
	params := req.URL.Query()
	params.Set("key", apiKey)
	params.Set("input", keyword)
	params.Set("inputtype", "textquery")
	params.Set("fields", "formatted_address,name,geometry")
	req.URL.RawQuery = params.Encode()
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to communicate with Google Geolocation API: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			println(err) // stderr
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %s", resp.Status)
	}
	content := struct {
		Candidates []struct {
			FormattedAddress string `json:"formatted_address"`
			Geometry         struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			} `json:"geometry"`
			Name string `json:"name"`
		} `json:"candidates"`
		Status string `json:"status"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, fmt.Errorf("failed to decode response of Google Geolocation API: %w", err)
	}
	coordinates := make([]Coordinate, len(content.Candidates))
	for i, v := range content.Candidates {
		coordinates[i] = Coordinate{
			Latitude:  v.Geometry.Location.Lat,
			Longitude: v.Geometry.Location.Lng,
			Label:     v.Name,
		}
	}
	return coordinates, nil
}
