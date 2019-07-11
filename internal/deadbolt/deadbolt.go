package deadbolt

import (
	"fmt"
	"net/http"
)

// Deadbolt struct represents data from the deadbolt.yml file
type Deadbolt struct {
	Path           string   // path/to/deadbolt.yml
	Port           string   `yaml:"port"`
	Secret         string   `yaml:"deadbolt_secret"`
	SSHDConfigPath string   `yaml:"sshd_config_path"`
	Whitelisted    []string `yaml:"whitelisted_clients"`
}

// New initializes a Deadbolt instance by loading deadbolt.yml and sets defaults.
// Environment variable overrides such as DEADBOLT_SECRET also take effect.
func New(path string) (*Deadbolt, error) {
	d := Deadbolt{Path: path}
	err := d.loadConfig()
	return &d, err
}

// Listen starts the deadbolt server.  The Deadbolt handler is responsible
// for validating requests are authorized.
func (dblt *Deadbolt) Listen() error {
	http.HandleFunc("/", dblt.defaultHandler)
	http.HandleFunc("/unlock", dblt.sshdHandler)
	http.HandleFunc("/lock", dblt.sshdHandler)

	fmt.Println("listening on port", dblt.Port)
	return http.ListenAndServe(":"+dblt.Port, logRequest(http.DefaultServeMux))
}

// PermitRootLogin updates the sshd_config with "PermitRootLogin <setting>".
// If /etc/ssh/sshd_config is changed, restart the sshd service.
func (dblt *Deadbolt) PermitRootLogin(setting string) error {
	sshd := sshdHandler{
		config:         dblt.SSHDConfigPath,
		restartAllowed: dblt.SSHDConfigPath == "/etc/ssh/sshd_config",
	}

	return sshd.PermitRootLogin(setting)
}
