package philote_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/fuzzingbits/philote"
)

func TestSuccess(t *testing.T) {
	site := &philote.Site{
		ContentPath: "./test_files/success",
	}

	targetTaxonomy := &philote.Taxonomy{
		Path:     "/",
		Markdown: "This is my homepage.",
		FrontMatter: &philote.FrontMatter{
			Title: "Homepage",
		},
		Children: []*philote.Taxonomy{
			{
				Path:     "/contact",
				Markdown: "This is my contact page.",
				FrontMatter: &philote.FrontMatter{
					Title: "Contact",
					Date:  time.Unix(1611144000, 0).UTC(),
				},
			},
			{
				Path:     "/about",
				Markdown: "This is my About Me page.",
				FrontMatter: &philote.FrontMatter{
					Title: "About",
					Date:  time.Unix(1610712000, 0).UTC(),
				},
			},
			{
				Path:     "/blog",
				Markdown: "This is my Blog.",
				FrontMatter: &philote.FrontMatter{
					Title: "Blog",
				},
				Children: []*philote.Taxonomy{
					{
						Path:     "/blog/first-post",
						Markdown: "This is my first post.",
						FrontMatter: &philote.FrontMatter{
							Title: "First Post",
						},
					},
				},
			},
		},
	}

	var targetErr error = nil

	testTaxonomy(t, site, targetTaxonomy, targetErr)
}

func testTaxonomy(t *testing.T, site *philote.Site, targetTaxonomy *philote.Taxonomy, targetError error) {
	if err := site.Prime(); err != nil {
		t.Fatalf("Error priming the site: %s", err.Error())
	}

	deepEqual(t, targetTaxonomy, site.Taxonomy)

}

func deepEqual(t *testing.T, expected interface{}, got interface{}) {
	expectedBytes, _ := json.Marshal(expected)
	gotBytes, _ := json.Marshal(got)

	if !reflect.DeepEqual(expectedBytes, gotBytes) {
		t.Fatalf("Got %s expected %s", string(gotBytes), string(expectedBytes))
	}
}
