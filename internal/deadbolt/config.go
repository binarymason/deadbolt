package deadbolt

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const defaultSSHDConfig = "/etc/ssh/sshd_config"
const defaultPort = "8080"

func (d *Deadbolt) setDefaults() {
	d.SSHDConfigPath = defaultSSHDConfig
	d.Port = defaultPort
}

func (d *Deadbolt) loadConfig() {
	yamlFile, err := ioutil.ReadFile(d.path)
	if err != nil {
		panic(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal: %v", err))
	}
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

func (d *Deadbolt) validate() {
	if d.Secret == "" {
		panic("Fatal: deadbolt secret not in environment or config file.")
	}
}
