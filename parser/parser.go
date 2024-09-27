package parser

import "io"

// Represents a HTML link
type Link struct {
	URL string
	Label string
}

// Takes an HTML doc and returns a slice of Links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
