package manifest

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	cases := []struct {
		FileName string
		Expect   *Manifest
	}{
		{
			FileName: "base-manifest.yml",
			Expect: &Manifest{
				Applications: []Application{
					Application{
						Name:        "app-name",
						Buildpack:   "",
						Command:     "",
						DiskQuota:   "",
						Memory:      "128M",
						Instances:   1,
						Domain:      "app.example.com",
						Domains:     []string(nil),
						Host:        "hello",
						Hosts:       []string(nil),
						NoHostName:  false,
						RandomRoute: false,
						NoRoute:     false,
						Path:        "path/to/app",
						Timeout:     0,
						Stack:       "",
						Env:         map[string]string(nil),
						Services:    []string(nil),
					},
				},
				Inherit: ""},
		},
	}

	for _, tc := range cases {
		t.Logf("Test ParseFile: %s", tc.FileName)

		path, err := filepath.Abs(filepath.Join("./fixtures", tc.FileName))
		if err != nil {
			t.Fatalf("Err: %s", err)
		}

		actual, err := ParseFile(path)
		if err != nil {
			t.Fatalf("Err: %s", err)
		}

		if !reflect.DeepEqual(actual, tc.Expect) {
			t.Fatalf("expect %#v to be eq %#v", actual, tc.Expect)
		}
	}
}
