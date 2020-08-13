package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly"
)

type GoReleasedVersion struct {
	FileName string `json:"filename"`
	Kind     string `json:"kind"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Size     string `json:"size"`
	Checksum string `json:"checksum"`
}

type GoAllreleasedVersion map[string]*GoReleasedVersion

func main() {

	allgo := make(GoAllreleasedVersion)

	jsonFile, _ := os.Create("./golang_version.json")
	defer jsonFile.Close()
	jsonWriter := io.Writer(jsonFile)

	c := colly.NewCollector(
		colly.AllowedDomains("golang.org"),
	)

	c.OnHTML("[id^=go]", func(e *colly.HTMLElement) {
		//if !validGoVersion.MatchString(e.Attr("id")) {
		//	return
		//}
		//		e.ForEach("table tbody",func(e *colly.HTMLElement){
		//      goReleaseVersion := GoReleasedVersion{}
		//      e.ForEach("tr",func(row *colly.HTMLElement){
		//
		//      })
		//		})
		fmt.Println("")
		fmt.Println("Go version: " + e.Attr("id"))
		fmt.Printf("%+v", e)
		allgo[e.Attr("id")] = &GoReleasedVersion{"https://golang.org/dl/go1.14.7.darwin-amd64.tar.gz", "Archive", "macOS", "x86-64", "119MB", "9a71abeb3de60ed33c0f90368be814d140bc868963e90fbb98ea665335ffbf9a"}
		fmt.Println("")
		spew.Dump(allgo)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://golang.org/dl/")

	encoder := json.NewEncoder(jsonWriter)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(&allgo)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}

}
