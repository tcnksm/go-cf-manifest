package manifest

// Manifest represent manifest file.
type Manifest struct {
	Applications []Application `yaml:"applications"`
	Inherit      string        `yaml:"inherit"`
}

// Application represents application attributes.
// https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html
type Application struct {
	Name        string            `yaml:"name"`
	Buildpack   string            `yaml:"buildpack"`
	Command     string            `yaml:"command"`
	DiskQuota   string            `yaml:"disk_quota"`
	Memory      string            `yaml:"memory"`
	Instances   int               `yaml:"instances"`
	Domain      string            `yaml:"domain"`
	Domains     []string          `yaml:"domains"`
	Host        string            `yaml:"host"`
	Hosts       []string          `yaml:"hosts"`
	NoHostName  bool              `yaml:"no-hostname"`
	RandomRoute bool              `yaml:"random-route"`
	NoRoute     bool              `yaml:"no-route"`
	Path        string            `yaml:"path"`
	Timeout     int               `yaml:"timeout"`
	Stack       string            `yaml:"stack"`
	Env         map[string]string `yaml:"env"`
	Services    []string          `yaml:"services"`
}
