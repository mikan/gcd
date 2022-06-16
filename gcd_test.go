package gcd

import "testing"

func TestHaversineDistance(t *testing.T) {
	hnd := Coordinate{Latitude: 35.5493932, Longitude: 139.7798386, Label: "羽田空港"}
	nrt := Coordinate{Latitude: 35.7719867, Longitude: 140.3928501, Label: "成田空港"}
	lhr := Coordinate{Latitude: 51.4700223, Longitude: -0.4542955, Label: "ヒースロー空港"}
	ord := Coordinate{Latitude: 41.98028, Longitude: -87.9089979, Label: "シカゴ・オヘア国際空港"}
	tests := []struct {
		from     Coordinate
		to       Coordinate
		expected float64
	}{
		{hnd, nrt, 60.66175074135226},
		{nrt, hnd, 60.66175074135226},
		{hnd, hnd, 0},
		{lhr, ord, 6344.451907097793},
	}
	for i, test := range tests {
		actual := HaversineDistance(test.from, test.to)
		if test.expected != actual {
			t.Errorf("#%d HaversineDistance(%s, %s) => %v (expected %v)", i, test.from.Label, test.to.Label, actual, test.expected)
		}
	}
}
