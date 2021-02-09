package philote

import (
	"bytes"
	"errors"
	"regexp"
)

var frontMatterFinder = regexp.MustCompile(`(?m)^---$`)

func parseMarkdown(fileBytes []byte) ([]byte, []byte, error) {
	matchGroups := frontMatterFinder.FindAllIndex(fileBytes, 2)

	if len(matchGroups) < 2 {
		return nil, nil, errors.New("didn't find markdown front matter")
	}

	frontMatterStart := matchGroups[0][1] + 1
	frontMatterStop := matchGroups[1][0] - 1
	bodyStart := matchGroups[1][1] + 1

	frontMatterBytes := fileBytes[frontMatterStart:frontMatterStop]
	bodyBytes := bytes.TrimSpace(fileBytes[bodyStart:])

	return frontMatterBytes, bodyBytes, nil
}
