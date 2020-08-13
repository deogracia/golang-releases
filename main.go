package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/gocolly/colly"
)

type GoReleasedVersion struct {
	Version  string `json:"version"`
	FileName string `json:"filename"`
	Kind     string `json:"kind"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Size     string `json:"size"`
	SHA256   string `json:"sha256"`
}

type GoAllreleasedVersion struct {
	GoReleasedVersion []GoReleasedVersion `json:"all"`
}

func main() {

	var allgo GoAllreleasedVersion
	var validGoVersion = regexp.MustCompile(`^go[1-9]((\.[0-9]+)*){2}$`)

	jsonFile, _ := os.OpenFile("./golang_version.json", os.O_CREATE, os.ModePerm)
	defer jsonFile.Close()

	c := colly.NewCollector(
		colly.AllowedDomains("golang.org"),
	)

	c.OnHTML("[id|=go]", func(e *colly.HTMLElement) {
		if !validGoVersion.MatchString(e.Attr("id")) {
			return
		}
		//    e.ForEach("table tbody",func(e *colly.HTMLElement){
		//      goReleaseVersion := GoReleasedVersion{}
		//      e.ForEach("tr",func(row *colly.HTMLElement){
		//
		//      })
		//    })
		fmt.Printf("%+v", e)
	})

	encoder := json.NewEncoder(jsonFile)
	encoder.Encode(allgo)
}
