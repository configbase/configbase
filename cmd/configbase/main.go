package main

import (
	"log"

	"github.com/BTBurke/configbase"
	"github.com/BTBurke/configbase/git"
	"github.com/BTBurke/configbase/kb"
)

func init() {
	checks := []func() error{
		kb.Init,
		git.Init,
	}
	for _, check := range checks {
		if err := check(); err != nil {
			log.Fatalf("startup: failure of init check: %v", err)
		}
	}
}

func main() {

	cmd := configbase.NewGitRunner("clone", "keybase://private/btburke/taxes")
	out, err := cmd.Run()
	if err != nil {
		log.Fatalf("git error: %s", err)
	}
	log.Printf("Success:\n%s", out)

	// cmd := exec.Command(gitpath, "clone", "keybase://private/btburke/taxes")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalf("git error: %v", err)
	// }
}
