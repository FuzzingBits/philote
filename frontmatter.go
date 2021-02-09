package philote

import "time"

// FrontMatter is a page
type FrontMatter struct {
	Title       string
	Date        time.Time
	Description string
}
