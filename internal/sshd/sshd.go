package sshd

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const SSHD_CONFIG = "/etc/ssh/sshd_config"

// Updates the sshd_config PermitRootLogin setting with argument "s"
// if /etc/ssh/sshd_config is changed, restart the sshd service.
func PermitRootLogin(s string) (err error) {
	if err := validateSetting(s); err != nil {
		return err
	}

	cfg, err := readConfig()

	if err != nil {
		return err
	}

	origMd5 := checksum(cfg)
	newConfig := generateConfig(s, string(cfg))
	newMd5 := checksum([]byte(newConfig))

	if err := writeConfig(newConfig); err != nil {
		return err
	}

	if newMd5 != origMd5 {
		err = restart()
	}

	return err
}

func validateSetting(s string) (err error) {
	v := false
	for _, x := range []string{"yes", "no", "without-password"} {
		if s == x {
			v = true
		}
	}
	if !v {
		err = errors.New("Invalid PermitRootLogin setting: " + s)
	}

	return err
}

func readConfig() ([]byte, error) {
	return ioutil.ReadFile(getConfigPath())
}

func getConfigPath() string {
	e := os.Getenv("DEADBOLT_SSHD_CONFIG")
	if e != "" {
		return e
	}
	return SSHD_CONFIG
}

func writeConfig(cfg string) error {
	p := getConfigPath()
	err := ioutil.WriteFile(p, []byte(cfg), 0644)
	return err
}

func generateConfig(m, cfg string) string {
	lines := strings.Split(cfg, "\n")
	for idx, line := range lines {
		setting := "PermitRootLogin"
		match, _ := regexp.Match(`^#?PermitRootLogin`, []byte(line))

		if match {
			lines[idx] = fmt.Sprintf("%s %s", setting, m)
		}
	}

	return strings.Join(lines, "\n")
}

func checksum(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Don't restart sshd agent unless ConfigPath is /etc/ssh/sshd_config (SSHD_CONFIG)
func restart() (err error) {
	if getConfigPath() == SSHD_CONFIG {
		cmd := exec.Command("service", "sshd", "restart")
		err = cmd.Run()
	}

	return err
}
