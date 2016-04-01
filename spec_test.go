package manifest

import (
	"reflect"
	"testing"
)

func TestManifest_Apply(t *testing.T) {
	cases := []struct {
		Manifest *Manifest
		Expect   *Manifest
	}{
		{
			&Manifest{
				Memory: "512M",
				Applications: []*Application{
					{
						Name:   "app1",
						Memory: "1G",
					},
					{
						Name: "app2",
					},
				},
			},

			&Manifest{
				Memory: "512M",
				Applications: []*Application{
					{
						Name:   "app1",
						Memory: "1G",
					},
					{
						Name:   "app2",
						Memory: "512M",
					},
				},
			},
		},

		{
			&Manifest{
				Instances: 2,
				Applications: []*Application{
					{
						Name: "app1",
						// Number will not be overwrited
						Instances: 3,
					},
					{
						Name: "app2",
						// Instances will be added
					},
				},
			},

			&Manifest{
				Instances: 2,
				Applications: []*Application{
					{
						Name:      "app1",
						Instances: 3,
					},
					{
						Name:      "app2",
						Instances: 2,
					},
				},
			},
		},

		// Merge slice
		{
			&Manifest{
				Services: []string{
					"redis",
					"mysql",
				},
				Applications: []*Application{
					{
						Name: "app1",
						Services: []string{
							"rabittmq",
						},
					},
					{
						Name: "app2",
					},
				},
			},

			&Manifest{
				Services: []string{
					"redis",
					"mysql",
				},
				Applications: []*Application{
					{
						Name: "app1",
						Services: []string{
							"rabittmq",
							"redis",
							"mysql",
						},
					},
					{
						Name: "app2",
						Services: []string{
							"redis",
							"mysql",
						},
					},
				},
			},
		},

		// Merge slice with same val
		{
			&Manifest{
				Services: []string{
					"redis",
				},
				Applications: []*Application{
					{
						Name: "app1",
						Services: []string{
							"redis",
							"rabbitmq",
						},
					},
				},
			},

			&Manifest{
				Services: []string{
					"redis",
				},
				Applications: []*Application{
					{
						Name: "app1",
						Services: []string{
							"redis",
							"rabbitmq",
						},
					},
				},
			},
		},

		// Merge map
		{
			&Manifest{
				Env: map[string]string{
					"RAILS_ENV": "test",
					"LOG_LEVEL": "INFO",
				},
				Applications: []*Application{
					{
						Name: "app1",
						Env: map[string]string{
							"RAILS_ENV": "production",
						},
					},
					{
						Name: "app2",
					},
				},
			},

			&Manifest{
				Env: map[string]string{
					"RAILS_ENV": "test",
					"LOG_LEVEL": "INFO",
				},
				Applications: []*Application{
					{
						Name: "app1",
						Env: map[string]string{
							"RAILS_ENV": "production",
							"LOG_LEVEL": "INFO",
						},
					},
					{
						Name: "app2",
						Env: map[string]string{
							"RAILS_ENV": "test",
							"LOG_LEVEL": "INFO",
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		err := tc.Manifest.Apply()
		if err != nil {
			t.Fatalf("Err: %s", err)
		}

		if !reflect.DeepEqual(tc.Manifest, tc.Expect) {
			t.Fatalf("expect \n %#v\n to be eq\n %#v", tc.Manifest.Applications[0], tc.Expect.Applications[0])
		}
	}

}
