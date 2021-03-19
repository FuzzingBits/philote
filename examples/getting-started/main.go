package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/fuzzingbits/philote"
)

func main() {
	// Create your instance of the philote.Site
	site := &philote.Site{
		Content:  os.DirFS("./content"),
		Template: template.Must(template.ParseFiles("./template.go.html")),
	}

	// Prime the site
	if err := site.Prime(); err != nil {
		panic(err)
	}

	// Serve the site
	http.ListenAndServe(":8090", site)
}
