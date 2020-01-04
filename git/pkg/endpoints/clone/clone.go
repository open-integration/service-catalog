package clone

import (
	"context"

	"github.com/open-integration/core/pkg/logger"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type (
	CloneOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *CloneArguments
	}
)

func Clone(opt CloneOptions) (*CloneReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	writer := log.FD()
	cloneOptions := &git.CloneOptions{
		URL:      opt.Arguments.Repo,
		Progress: writer,
		Auth:     buildAuthMethod(opt.Arguments.Provider, nil),
	}
	_, err := git.PlainClone(opt.Arguments.Path, false, cloneOptions)
	if err != nil {
		return nil, err
	}
	return &CloneReturns{
		Location: opt.Arguments.Path,
	}, nil
}

func buildAuthMethod(provider Provider, auth *GithubAuth) transport.AuthMethod {
	if auth.Token != nil {
		return &http.BasicAuth{
			Username: string(provider),
			Password: *auth.Token,
		}
	}
	return &http.BasicAuth{
		Username: *auth.Username,
		Password: *auth.Password,
	}
}
