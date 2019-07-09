package validate

import "github.com/binarymason/go-deadbolt/internal/config"

func validRequest(ip, auth string, cfg config.Config) bool {
	return validIP(ip, cfg.Whitelisted) && validAuth(auth, cfg.Secret)
}

func validIP(ip string, whitelisted []string) bool {
	for _, w := range whitelisted {
		if ip == w {
			return true
		}
	}

	return false
}

func validAuth(a, s string) bool {
	return a == s
}
