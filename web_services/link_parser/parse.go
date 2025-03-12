package linkparser

import "io"

type LinkParserOutput struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]LinkParserOutput, error) {
	return nil, nil
}
