package kb

import (
	"os/exec"
)

func Init() error {
	if _, err := exec.LookPath("keybase"); err != nil {
		return err
	}
	if _, err := exec.LookPath("git-remote-keybase"); err != nil {
		return err
	}
	return nil
}
