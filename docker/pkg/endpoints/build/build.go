package build

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
	"github.com/open-integration/core/pkg/logger"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type (
	BuildOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *BuildArguments
	}
)

func Build(opt BuildOptions) (*BuildReturns, error) {
	log := logger.NewFromFilePath(opt.LoggerFD)
	args := opt.Arguments
	log.Debug("Connecting to daemon", "host", args.Host, "api-version", args.APIVersion, "build-context", args.BuildContext)
	c, err := client.NewClient(args.Host, args.APIVersion, nil, nil)
	if err != nil {
		return nil, err
	}
	buildOptions := types.ImageBuildOptions{
		Tags:       []string{args.Tag},
		Dockerfile: args.Dockerfile,
	}
	dest := path.Join("/tmp", fmt.Sprintf("%s.tar", stringWithCharset(10)))
	if err := buildTar(args.BuildContext, dest); err != nil {
		return nil, err
	}
	log.Debug("Created tar", "name", dest)

	buildContext, err := os.Open(dest)
	if err != nil {
		return nil, err
	}
	defer buildContext.Close()
	resp, err := c.ImageBuild(opt.Context, buildContext, buildOptions)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Debug("Build finished")
	log.Debug(string(data))
	defer resp.Body.Close()
	os.RemoveAll(dest)
	return &BuildReturns{}, nil
}

func buildTar(source string, destination string) error {
	tar := new(archivex.TarFile)
	if err := tar.Create(destination); err != nil {
		return err
	}

	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if err := tar.AddAll(path, false); err != nil {
				return err
			}
		} else {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			if err := tar.Add(info.Name(), f, info); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return tar.Close()
}

func stringWithCharset(length int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
