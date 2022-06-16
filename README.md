gcd
===

A golang-implemented GCD (great-circular distance) calculator with geolocation finding feature using Google Geolocation API.

Go で実装した大円距離計算機です。Google Geolocation API で地点を検索する機能もあります。

## Library

```
go get github.com/mikan/gcd
```

GCD calculation example:

```go
package main

import "github.com/mikan/gcd"

func main() {
	from := gcd.Coordinate{Latitude: 35.7719867, Longitude: 140.3928501}
	to := gcd.Coordinate{Latitude: 35.76483331, Longitude: 140.3860192}
	gcd.HaversineDistance(from, to) // 60.66175074135226 (km)
}
```

Location finding example:

```go
package main

import "github.com/mikan/gcd"

func main() {
	from, err := gcd.FindPlaceWithGoogle("API_KEY", "Haneda Airport")
	if err != nil {
		panic(err)
	}
	to, err := gcd.FindPlaceWithGoogle("API_KEY", "Narita Airport")
	if err != nil {
		panic(err)
	}
	if len(from) > 0 && len(to) > 0 {
		gcd.HaversineDistance(from[0], to[0]) // 60.66175074135226 (km)
	}
}
```

## CLI

```
go install github.com/mikan/gcd@latest
```

Usage:

```
gcd -key <API_KEY> <FROM_KEYWORD> <TO_KEYWORD>
```

Example:

```
% gcd -key xxx "Haneda Airport" "Narita Airport"
Haneda Airport:
* 35.5493932,139.7798386 羽田空港
Narita Airport:
* 35.7719867,140.3928501 成田国際空港
60.66175074135226km
```

## Author

- [mikan](https://github.com/mikan)

## License

[BSD 3-clause License](LICENSE)
