package philote

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
)

// Taxonomy is the overall structure of the page
type Taxonomy struct {
	FrontMatter *FrontMatter
	Markdown    string
	Path        string
	Children    []*Taxonomy
	Parent      *Taxonomy `json:"-"`
}

// Render the markdown into HTML
func (taxonomy *Taxonomy) Render() template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(taxonomy.Markdown)))
}

func (taxonomy *Taxonomy) readInContent(reader io.Reader) error {
	fileBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	frontMatterBytes, markdownBytes, err := parseMarkdown(fileBytes)
	if err != nil {
		return errors.New("unable parse markdown file")
	}

	content := &FrontMatter{}
	if err := yaml.Unmarshal(frontMatterBytes, &content); err != nil {
		return err
	}

	taxonomy.FrontMatter = content
	taxonomy.Markdown = string(markdownBytes)

	return nil
}

func buildTaxonomy(startingDirectory string, startingPath string) (*Taxonomy, error) {
	files, err := ioutil.ReadDir(startingDirectory)
	if err != nil {
		return nil, err
	}

	taxonomy := &Taxonomy{
		Path: strings.TrimSuffix(startingPath, "/"),
	}

	if taxonomy.Path == "" {
		taxonomy.Path = "/"
	}

	for _, f := range files {
		fileName := f.Name()
		filePath := startingDirectory + "/" + fileName
		fullPath := startingPath + strings.TrimSuffix(fileName, ".md")
		file, err := os.Open(startingDirectory + "/" + fileName)
		if err != nil {
			return nil, err
		}

		defer file.Close()

		if fileName == "_index.md" {
			if err := taxonomy.readInContent(file); err != nil {
				return nil, err
			}

			continue
		}

		if f.IsDir() {
			subTaxonomy, err := buildTaxonomy(filePath, fullPath+"/")
			if err != nil {
				return nil, err
			}

			subTaxonomy.Parent = taxonomy

			taxonomy.Children = append(taxonomy.Children, subTaxonomy)
			continue
		}

		subTaxonomy := &Taxonomy{
			Path:   fullPath,
			Parent: taxonomy,
		}
		if err := subTaxonomy.readInContent(file); err != nil {
			return nil, err
		}
		taxonomy.Children = append(taxonomy.Children, subTaxonomy)
	}

	sort.Slice(taxonomy.Children, func(p, q int) bool {
		return taxonomy.Children[p].FrontMatter.Date.After(taxonomy.Children[q].FrontMatter.Date)
	})

	return taxonomy, nil
}
