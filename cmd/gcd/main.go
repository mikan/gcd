package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/mikan/gcd"
)

func main() {
	// Usage: gcd -key <API_KEY> <FROM_KEYWORD> <TO_KEYWORD>
	// Example: gcd -key xxx "Narita Airport" "Kansai Airport"
	apiKey := flag.String("key", "", "API key of Google Maps API")
	flag.Parse()
	if len(*apiKey) == 0 {
		println("Please specify the API key with -key option.")
		flag.Usage()
		os.Exit(2)
	}
	if flag.NArg() != 2 {
		println("Please specify two location keywords")
		os.Exit(2)
	}

	// find "from" coordinate
	from, err := find(*apiKey, flag.Arg(0))
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	// find "to" coordinate
	to, err := find(*apiKey, flag.Arg(1))
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%vkm\n", gcd.HaversineDistance(*from, *to))
}

func find(apiKey, keyword string) (*gcd.Coordinate, error) {
	candidates, err := gcd.FindPlaceWithGoogle(apiKey, keyword)
	if err != nil {
		return nil, err
	}
	if len(candidates) == 0 {
		return nil, errors.New("No match for \"" + keyword + "\"")
	}
	fmt.Printf("%s:\n", keyword)
	for i, c := range candidates {
		if i == 0 {
			fmt.Printf("* %s\n", c.String())
		} else {
			fmt.Printf("  %s\n", c.String())
		}
	}
	return &candidates[0], nil
}
