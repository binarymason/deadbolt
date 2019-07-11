package deadbolt

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

type sshdHandler struct {
	config         string
	restartAllowed bool
}

// PermitRootLogin updates the sshd_config with "PermitRootLogin <setting>".
// If /etc/ssh/sshd_config is changed, restart the sshd service.
func (sshd *sshdHandler) PermitRootLogin(setting string) (err error) {
	if err := validateSetting(setting); err != nil {
		return err
	}

	cfg, err := sshd.read()

	if err != nil {
		return err
	}

	origMd5 := md5sum(cfg)
	newConfig := generateConfig(setting, cfg)
	newMd5 := md5sum(newConfig)

	if err := sshd.write(newConfig); err != nil {
		return err
	}

	if newMd5 != origMd5 {
		return sshd.restart()
	}

	return nil
}

func (sshd *sshdHandler) restart() (err error) {
	if !sshd.restartAllowed {
		return nil
	}

	cmd := exec.Command("service", "sshd", "restart")
	return cmd.Run()
}

func validateSetting(s string) (err error) {
	valid := false
	for _, x := range []string{"yes", "no", "without-password"} {
		if s == x {
			valid = true
		}
	}
	if !valid {
		err = errors.New("Invalid PermitRootLogin setting: " + s)
	}

	return err
}

func (sshd *sshdHandler) read() ([]byte, error) {
	return ioutil.ReadFile(sshd.config)
}

func (sshd *sshdHandler) write(c []byte) error {
	return ioutil.WriteFile(sshd.config, c, 0644)
}

func generateConfig(m string, c []byte) []byte {
	cfg := string(c)
	lines := strings.Split(cfg, "\n")
	for idx, line := range lines {
		setting := "PermitRootLogin"
		match, _ := regexp.Match(`^#?PermitRootLogin`, []byte(line))

		if match {
			lines[idx] = fmt.Sprintf("%s %s", setting, m)
		}
	}

	return []byte(strings.Join(lines, "\n"))
}

func md5sum(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
