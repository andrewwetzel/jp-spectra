package main

import (
	"fmt"
	"log"

	"github.com/andrewwetzel/jp-spectra/fetcher"
)

func main() {
	url := "https://opus.pds-rings.seti.org/api/data.json?target=Jupiter&instrument=Cassini+VIMS&limit=3"
	data, err := fetcher.FetchJSON(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
