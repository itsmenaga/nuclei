package templates

import (
	"os"

	"gopkg.in/yaml.v2"
)

// ParseTemplate parses a yaml request template file
func ParseTemplate(file string) (*Template, error) {
	template := &Template{}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	err = yaml.NewDecoder(f).Decode(template)
	if err != nil {
		f.Close()
		return nil, err
	}
	f.Close()

	// Compile the matchers and the extractors
	for _, request := range template.Requests {
		for _, matcher := range request.Matchers {
			if err = matcher.CompileMatchers(); err != nil {
				return nil, err
			}
		}

		for _, extractor := range request.Extractors {
			if err := extractor.CompileExtractors(); err != nil {
				return nil, err
			}
		}
	}
	return template, nil
}
