package configbase

import (
	"context"
	"errors"
	"os/exec"
)

// Runner is a generic structure to manage shelled out commands with
// optional context
type Runner struct {
	Cmd []string
	Ctx context.Context
}

// NewRunner returns a Runner with a background context
func NewRunner(cmd ...string) *Runner {
	return &Runner{
		Cmd: cmd,
		Ctx: context.Background(),
	}
}

// NewContextRunner returns a runner with a given context
func NewContextRunner(ctx context.Context, cmd ...string) *Runner {
	return &Runner{
		Cmd: cmd,
		Ctx: ctx,
	}
}

// NewKeybaseRunner returns a runner configured to run the command
// using the keybase client with JSON output
func NewKeybaseRunner(cmd ...string) *Runner {
	return &Runner{
		Cmd: append([]string{"keybase", "--json"}, cmd...),
		Ctx: context.Background(),
	}
}

// NewKeybaseContextRunner returns a runner configured to run the command
// using the keybase client with JSON output and custom context
func NewKeybaseContextRunner(ctx context.Context, cmd ...string) *Runner {
	return &Runner{
		Cmd: append([]string{"keybase", "--json"}, cmd...),
		Ctx: ctx,
	}
}

// NewGitRunner returns a runner configured to run the command using
// the git client
func NewGitRunner(cmd ...string) *Runner {
	return &Runner{
		Cmd: append([]string{"git"}, cmd...),
		Ctx: context.Background(),
	}
}

// NewGitContextRunner returns a runner configured to run the command using
// the git client and custom context
func NewGitContextRunner(ctx context.Context, cmd ...string) *Runner {
	return &Runner{
		Cmd: append([]string{"git"}, cmd...),
		Ctx: ctx,
	}
}

// Run will run the command and return combined Stderr and Stdout as byte slice
// or error
func (r *Runner) Run() ([]byte, error) {

	var cmd *exec.Cmd
	switch len(r.Cmd) {
	case 0:
		return nil, errors.New("no command to run")
	case 1:
		cmd = exec.CommandContext(r.Ctx, r.Cmd[0])
	default:
		cmd = exec.CommandContext(r.Ctx, r.Cmd[0], r.Cmd[1:]...)
	}
	return cmd.CombinedOutput()
}
