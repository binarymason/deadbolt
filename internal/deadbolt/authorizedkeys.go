package deadbolt

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// writeAuthorizedKeys creates a completely new authorized_keys file every time
// with whatever authorized keys are specified to deadbolt.
//
// TODO: handle any pre-existing keys instead of deleting them.
func (d *Deadbolt) writeAuthorizedKeys() error {
	if len(d.AuthorizedKeys) == 0 {
		return errors.New("No authorized keys given to deadbolt")
	}

	mkdirP(d.AuthorizedKeysFile)

	log.Println("writing authorized keys to ", d.AuthorizedKeysFile)
	keys := []byte(strings.Join(d.AuthorizedKeys, "\n"))
	return ioutil.WriteFile(d.AuthorizedKeysFile, keys, 0600)
}

func mkdirP(p string) {
	absPath, _ := filepath.Abs(p)
	dir := filepath.Dir(absPath)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}
}
