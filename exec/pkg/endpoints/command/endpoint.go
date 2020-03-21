package command

import (
	"context"
	"os/exec"

	"github.com/open-integration/core/pkg/logger"
)

type (
	CommandOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *CommandArguments
	}
)

func Command(opt CommandOptions) (*CommandReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	cmd := exec.Command("/bin/sh", "-c", opt.Arguments.Command)
	cmd.Stdout = log.FD()
	cmd.Stderr = log.FD()
	err := cmd.Run()
	return &CommandReturns{}, err
}
