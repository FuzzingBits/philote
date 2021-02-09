package philote

import (
	"bytes"
	"html/template"
	"net/http"
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
	taxonomy, found := site.pathMap[r.URL.Path]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("philote not found"))
		return
	}

	buffer := bytes.NewBuffer([]byte{})
	err := site.Template.Execute(buffer, TemplatePayload{
		Taxonomy: taxonomy,
	})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
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
