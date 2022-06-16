package gcd

import (
	"fmt"
	"math"
	"strings"
)

const earthRadiusKM = 6371.0710 // mean earth radius as kilometers

// Coordinate represents combination of latitude and longitude.
type Coordinate struct {
	Latitude  float64
	Longitude float64
	Label     string
}

func (c *Coordinate) String() string {
	return strings.TrimSpace(fmt.Sprintf("%v,%v %s", c.Latitude, c.Longitude, c.Label))
}

// HaversineDistance calculates great-circular distance as kilometers between two Coordinate using haversine formula.
func HaversineDistance(from, to Coordinate) float64 {
	var fromR = from.Latitude * (math.Pi / 180)
	var toR = to.Latitude * (math.Pi / 180)
	var diffLat = toR - fromR
	var diffLon = (to.Longitude - from.Longitude) * (math.Pi / 180)
	return 2 * earthRadiusKM * math.Asin(math.Sqrt(
		math.Sin(diffLat/2)*math.Sin(diffLat/2)+
			math.Cos(fromR)*math.Cos(toR)*math.Sin(diffLon/2)*math.Sin(diffLon/2)))
}
