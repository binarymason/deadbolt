package sshd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const SSHD_CONFIG = "/etc/ssh/sshd_config"

func PermitRootLogin(s string) (err error) {
	if err := validateSetting(s); err != nil {
		return err
	}
	cfg, err := readConfig()

	if err != nil {
		return err
	}
	err = writeConfig(generateConfig(s, string(cfg)))
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
