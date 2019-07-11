package deadbolt

import (
	"fmt"
	"net/http"
)

type Deadbolt struct {
	path           string
	Port           string   `yaml:"port"`
	Secret         string   `yaml:"deadbolt_secret"`
	SSHDConfigPath string   `yaml:"sshd_config_path"`
	Whitelisted    []string `yaml:"whitelisted_clients"`
}

func New(path string) *Deadbolt {
	d := Deadbolt{path: path}
	d.setDefaults()
	d.loadConfig()
	d.setOverrides()
	d.validate()
	return &d
}

func (dblt *Deadbolt) Listen() error {
	http.HandleFunc("/", dblt.defaultHandler)
	http.HandleFunc("/unlock", dblt.sshdHandler)
	http.HandleFunc("/lock", dblt.sshdHandler)

	fmt.Println("listening on port", dblt.Port)
	return http.ListenAndServe(":"+dblt.Port, logRequest(http.DefaultServeMux))
}

func (dblt *Deadbolt) PermitRootLogin(setting string) error {
	sshd := sshdHandler{
		config:         dblt.SSHDConfigPath,
		restartAllowed: dblt.SSHDConfigPath == "/etc/ssh/sshd_config",
	}
	return sshd.PermitRootLogin(setting)
}
