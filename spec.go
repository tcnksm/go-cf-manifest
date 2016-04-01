package manifest

import (
	"fmt"
	"reflect"
)

// Manifest represent manifest file.
// See more details on https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html
type Manifest struct {
	Name         string            `yaml:"name"`
	Buildpack    string            `yaml:"buildpack"`
	Command      string            `yaml:"command"`
	DiskQuota    string            `yaml:"disk_quota"`
	Domain       string            `yaml:"domain"`
	Domains      []string          `yaml:"domains"`
	Stack        string            `yaml:"stack"`
	Instances    int               `yaml:"instances"`
	Memory       string            `yaml:"memory"`
	Host         string            `yaml:"host"`
	Hosts        []string          `yaml:"hosts"`
	NoHostName   bool              `yaml:"no-hostname"`
	RandomRoute  bool              `yaml:"random-route"`
	Path         string            `yaml:"path"`
	Timeout      int               `yaml:"timeout"`
	NoRoute      bool              `yaml:"no-route"`
	Env          map[string]string `yaml:"env"`
	Services     []string          `yaml:"services"`
	Applications []*Application    `yaml:"applications"`
	Inherit      string            `yaml:"inherit"`
}

// Application represents application configuration
type Application struct {
	Name        string            `yaml:"name"`
	Buildpack   string            `yaml:"buildpack"`
	Command     string            `yaml:"command"`
	DiskQuota   string            `yaml:"disk_quota"`
	Domain      string            `yaml:"domain"`
	Domains     []string          `yaml:"domains"`
	Stack       string            `yaml:"stack"`
	Instances   int               `yaml:"instances"`
	Memory      string            `yaml:"memory"`
	Host        string            `yaml:"host"`
	Hosts       []string          `yaml:"hosts"`
	NoHostName  bool              `yaml:"no-hostname"`
	RandomRoute bool              `yaml:"random-route"`
	Path        string            `yaml:"path"`
	Timeout     int               `yaml:"timeout"`
	NoRoute     bool              `yaml:"no-route"`
	Env         map[string]string `yaml:"env"`
	Services    []string          `yaml:"services"`
}

// Apply applies above the application block (common configuration)
// to applications block.
//
// Rule is that content in the applications block overrides
// content above the applications block, if the two conflict.
//
// (TODO:tcnksm) Need more refactoring
func (m *Manifest) Apply() error {

	if len(m.Applications) == 0 {
		return fmt.Errorf("no applications field found")
	}

	// Find the non-zero value from above the applications block.
	// If the non-zero value is found, access that filed on applications
	// block and set it if it's empty.
	//
	// If filed is slice or map then append value (TODO: Check it).
	elem := reflect.ValueOf(m).Elem()
	for i := 0; i < elem.NumField(); i++ {

		field := elem.Type().Field(i).Name

		// Ignore
		if field == "Inherit" || field == "Applications" {
			continue
		}

		val := elem.Field(i)
		switch val.Kind() {
		case reflect.Slice, reflect.Map:
			if val.Len() == 0 {
				continue
			}
		default:
			zeroVal := reflect.Zero(val.Type())
			if val.Interface() == zeroVal.Interface() {
				continue
			}
		}

		for _, app := range m.Applications {

			appVal := reflect.ValueOf(app)
			appElem := appVal.Elem()
			if appElem.Kind() != reflect.Struct {
				continue
			}

			appField := appElem.FieldByName(field)
			if appField.Kind() == reflect.String {
				if appField.Interface() != "" {
					continue
				}
				x, ok := val.Interface().(string)
				if ok {
					appField.SetString(x)
				}
			}

			if appField.Kind() == reflect.Int {
				if appField.Interface() != 0 {
					continue
				}

				x, ok := val.Interface().(int)
				if ok {
					appField.SetInt(int64(x))
				}
			}

			if appField.Kind() == reflect.Slice {
				c, _ := val.Interface().([]string)
				a, _ := appField.Interface().([]string)
				var targets []string
				for _, v1 := range c {
					f := false
					for _, v2 := range a {
						if v1 == v2 {
							f = true
						}
					}

					if !f {
						targets = append(targets, v1)
					}
				}

				a = append(a, targets...)
				appField.Set(reflect.ValueOf(a))
			}

			if appField.Kind() == reflect.Map {
				newMap := make(map[string]string)
				c, _ := val.Interface().(map[string]string)
				for k, v := range c {
					newMap[k] = v
				}

				a, _ := appField.Interface().(map[string]string)
				for k, v := range a {
					newMap[k] = v
				}
				appField.Set(reflect.ValueOf(newMap))
			}
		}
	}

	return nil
}

// Merge merges the manifest. the given other will overwrite
// target manifest. This fucntion is supposed to be used
// for merging inherit manifest.
//
// TODO
func (m *Manifest) merge(other *Manifest) error {
	return nil
}

// Validate validates manifest has requirement filed or not.
//
// TODO
func (m *Manifest) validate() error {
	return nil
}
