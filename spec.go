package manifest

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
	Applications []Application     `yaml:"applications"`
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
