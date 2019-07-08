package sshd

import (
	"regexp"
	"strings"
)

func lockConfig(cfg string) string {
	return generateConfig("lock", cfg)
}

func unlockConfig(cfg string) string {
	return generateConfig("unlock", cfg)
}

func generateConfig(m, cfg string) string {
	lines := strings.Split(cfg, "\n")
	var result []string
	for _, line := range lines {
		result = append(result, toggle(m, line))
	}

	return strings.Join(result, "\n")
}

func toggle(m, s string) string {
	setting := "PermitRootLogin"
	match, _ := regexp.Match(`PermitRootLogin`, []byte(s))

	if !match {
		return s
	}

	switch m {
	case "lock":
		return setting + " no"
	case "unlock":
		return setting + " yes"
	default:
		panic("unhandled toggle method: " + m)

	}
}
