package philote

import (
	"bytes"
	"html/template"
	"net/http"
	"time"

	"github.com/fuzzingbits/forge"
)

// Site is a philote site
type Site struct {
	ContentPath string
	Template    *template.Template
	Taxonomy    *Taxonomy
	pathMap     map[string]*Taxonomy
}

// ServeHTTP the site based on the path provided
func (site *Site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	taxonomy, found := site.pathMap[r.URL.Path]
	if !found {
		statusCode = http.StatusNotFound
		taxonomy = &Taxonomy{
			FrontMatter: &FrontMatter{
				Title: "404 Not Found",
				Date:  time.Now(),
			},
			Markdown: "The requested page was not found.",
		}
	}

	buffer := bytes.NewBuffer([]byte{})
	err := site.Template.Execute(buffer, TemplatePayload{
		Site:     site,
		Taxonomy: taxonomy,
	})
	if err != nil {
		panic(err)
	}

	forge.RespondHTML(w, statusCode, buffer.Bytes())
}

// Prime the site
func (site *Site) Prime() error {
	var err error
	site.Taxonomy, err = buildTaxonomy(site.ContentPath, "/")

	site.pathMap = make(map[string]*Taxonomy)
	site.mapTaxonomy(site.Taxonomy)

	return err
}

func (site *Site) mapTaxonomy(taxonomy *Taxonomy) {
	site.pathMap[taxonomy.Path] = taxonomy
	for _, childTaxonomy := range taxonomy.Children {
		site.mapTaxonomy(childTaxonomy)
	}
}
