package main

import "fmt"

func newPublication(pubType string, name string, pages int, publisher string) (iPublication, error){

	if pubType == "newspaper" {
		return createNewspaper(name, pages, publisher), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pages, publisher), nil
	}

	return nil, fmt.Errorf("unknown publication type")
}
