package git

import "os/exec"

func Init() error {
	_, err := exec.LookPath("git")
	return err
}
