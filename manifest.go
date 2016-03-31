/*
go-cf-manifest is pacakge for handling CloudFoundry `manifest.yml` file.

  m, _ := manifest.ParseFile("./manifest.yml")
  fmt.Printf("%#v",m)

*/
package manifest

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Parse parses manifest file given as io.Reader.
func Parse(rd io.Reader) (*Manifest, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, rd); err != nil {
		return nil, err
	}

	var manifest Manifest
	if err := yaml.Unmarshal(buf.Bytes(), &manifest); err != nil {
		return nil, err
	}

	return &manifest, nil
}

// ParseFile parses a manifest file.
func ParseFile(path string) (*Manifest, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Parse(f)
}
