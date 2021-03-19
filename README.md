# philote

Minimalist library to serve markdown content as webpages.

[![Documentation](https://godoc.org/github.com/fuzzingbits/philote?status.svg)](https://pkg.go.dev/github.com/fuzzingbits/philote)
[![GitHub Actions](https://github.com/fuzzingbits/philote/workflows/Main/badge.svg)](https://github.com/fuzzingbits/philote/actions)
[![Coverage Status](https://coveralls.io/repos/github/fuzzingbits/philote/badge.svg?branch=main)](https://coveralls.io/github/fuzzingbits/philote?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/fuzzingbits/philote)](https://goreportcard.com/report/github.com/fuzzingbits/philote)
[![License](https://img.shields.io/github/license/fuzzingbits/philote)](https://github.com/fuzzingbits/philote/blob/main/LICENSE)

# Index
- [Disclaimer](#disclaimer)
- [Install](#install)
- [Getting Started](#getting-started)
- [Templates](#templates)
- [Taxonomy](#taxonomy)
- [FrontMatter](#frontmatter)

# Disclaimer

This library is under active development and subject to breaking changes at any time.

# Install
```
go get github.com/fuzzingbits/philote
```

# Getting Started

The following steps are how the [Getting Started Example](examples/getting-started) was created.

1. Serve a philote.Site
```go
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
```
2. Create the Go HTML Template (`./template.go.html`):
```html
<!doctype html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Philote Site</title>
	</head>
	<body>
		<h1>Title: {{.Taxonomy.FrontMatter.Title}}</h1>
		<p>Description: {{.Taxonomy.FrontMatter.Description}}</p>
		{{.Taxonomy.Render}}
	</body>
</html>
```
3. Create your first page (`.content/_index.md`):
```md
---
title: My First Page
date: 2021-01-01T00:00:00Z
description: "This is my first page"
---

This is being served with [philote](https://github.com/fuzzingbits/philote).
```

# Templates

- Philote uses Go Standard Library HTML templates so all of the [html/template documentation](https://golang.org/pkg/html/template/) applies.
- A single struct is always passed into the template you provide: [TemplatePayload](https://pkg.go.dev/github.com/fuzzingbits/philote#TemplatePayload).

# Taxonomy

The Site Taxonomy is automatically derived from the contents of the Content FileSystem:
- all files are used as pages with their path and base filename used as the URL
	- Example: `./posts/hello-world.md` == `http://localhost:8000/posts/hello-world`
- all files must be have the `.md` extension
- every directory (including the root directory) must include a `_index.md` file.

Example:
```
┌── _index.md
├── about.md
└── posts
    ├── _index.md
    └── hello-world.md
```

# FrontMatter

Every Markdown file must include FrontMatter at the beginning of the file. The FrontMatter should be in YAML format with fields matching the [FrontMatter](https://pkg.go.dev/github.com/fuzzingbits/philote#FrontMatter) struct.

Example:
```md
---
title: My First Page
date: 2021-01-01T00:00:00Z
description: "This is my first page"
---

This is being served with [philote](https://github.com/fuzzingbits/philote).
```
