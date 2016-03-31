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
			"base.yml",
			&Manifest{
				Name:      "sample",
				Buildpack: "https://github.com/cloudfoundry/go-buildpack",
				Command:   "bundle exec rake VERBOSE=true",
				DiskQuota: "1024M",
				Domain:    "sample.example.com",
				Domains: []string{
					"example1.com",
					"example2.org",
				},
				Stack:     "cflinuxfs2",
				Instances: 2,
				Memory:    "128M",
				Host:      "hello",
				Hosts: []string{
					"hello1",
					"hello2",
				},
				NoHostName:  true,
				RandomRoute: true,
				Path:        "path/to/app",
				Timeout:     80,
				NoRoute:     true,
				Env: map[string]string{
					"RAILS_ENV": "production",
					"RACK_ENV":  "production",
				},
				Services: []string{
					"mysql",
					"redis",
				},

				Applications: []Application(nil),
				Inherit:      "",
			},
		},

		{
			"applications.yml",
			&Manifest{
				Applications: []Application{
					Application{
						Name:      "sample",
						Buildpack: "https://github.com/cloudfoundry/go-buildpack",
						Command:   "bundle exec rake VERBOSE=true",
						DiskQuota: "1024M",
						Domain:    "sample.example.com",
						Domains: []string{
							"example1.com",
							"example2.org",
						},
						Stack:     "cflinuxfs2",
						Instances: 2,
						Memory:    "128M",
						Host:      "hello",
						Hosts: []string{
							"hello1",
							"hello2",
						},
						NoHostName:  true,
						RandomRoute: true,
						Path:        "path/to/app",
						Timeout:     80,
						NoRoute:     true,
						Env: map[string]string{
							"RAILS_ENV": "production",
							"RACK_ENV":  "production",
						},
						Services: []string{
							"mysql",
							"redis",
						},
					},
				},
				Inherit: "",
			},
		},

		{
			"empty.yml",
			&Manifest{
				Applications: []Application(nil),
				Inherit:      "",
			},
		},
	}

	for _, tc := range cases {
		t.Logf("TestParseFile: %s", tc.FileName)

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
