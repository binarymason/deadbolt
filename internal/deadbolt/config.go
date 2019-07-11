package deadbolt

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func (d *Deadbolt) loadConfig() error {
	d.setDefaults()

	if err := d.unmarshalConfig(); err != nil {
		return err
	}

	d.setOverrides()
	return d.validate()
}

func (d *Deadbolt) setDefaults() {
	d.AuthorizedKeysFile = os.Getenv("HOME") + "/.ssh/authorized_keys"
	d.Port = "8080"
	d.SSHDConfigPath = "/etc/ssh/sshd_config"
}

func (d *Deadbolt) unmarshalConfig() error {
	yamlFile, err := ioutil.ReadFile(d.Path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlFile, &d)
}

func (d *Deadbolt) setOverrides() {
	s := os.Getenv("DEADBOLT_SECRET")
	if s != "" {
		d.Secret = s
	}

	sshdCfg := os.Getenv("DEADBOLT_SSHD_CONFIG")
	if sshdCfg != "" {
		d.SSHDConfigPath = sshdCfg
	}
}

func (d *Deadbolt) validate() error {
	if d.Secret == "" {
		return errors.New("deadbolt secret not in environment or config file")
	}

	if fileNotFound(d.SSHDConfigPath) {
		return errors.New("ssh config file does not exist: " + d.SSHDConfigPath)
	}

	return nil
}

// fileNotFound returns true if a file path does not exist or is a directory.
func fileNotFound(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return true
	}
	return info.IsDir()
}
