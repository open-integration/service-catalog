package run

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/open-integration/core/pkg/logger"
)

type (
	RunOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *RunArguments
	}

	runner interface {
		ContainerCreate(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, string) (container.ContainerCreateCreatedBody, error)
		ContainerStart(context.Context, string, types.ContainerStartOptions) error
		ContainerWait(context.Context, string) (int64, error)
		ContainerLogs(context.Context, string, types.ContainerLogsOptions) (io.ReadCloser, error)
	}

	runOptions struct {
		logerWriter                     io.WriteCloser
		containerCreateConfig           *container.Config
		containerCreateHostConfig       *container.HostConfig
		containerCreateNetworkingConfig *network.NetworkingConfig
		containerStartConfig            *types.ContainerStartOptions
		containerLogsConfig             *types.ContainerLogsOptions
	}
)

func Run(opt RunOptions) (*RunReturns, error) {
	log := logger.NewFromFilePath(opt.LoggerFD)
	writer, err := os.OpenFile(opt.LoggerFD, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	defer writer.Close()
	args := opt.Arguments
	log.Debug("Connecting to daemon", "host", args.Host, "api-version", args.APIVersion)
	c, err := client.NewClient(args.Host, args.APIVersion, nil, nil)
	if err != nil {
		return nil, err
	}
	cnf, err := buildRunConfig(&opt)
	if err != nil {
		return nil, err
	}
	defer cnf.logerWriter.Close()
	err = run(opt.Context, log, c, cnf)
	return &RunReturns{}, err
}

func run(context context.Context, logger logger.Logger, runner runner, runOptions *runOptions) error {
	logger.Debug("Creating container", "config", runOptions.containerCreateConfig)
	respBody, err := runner.ContainerCreate(context, runOptions.containerCreateConfig, runOptions.containerCreateHostConfig, runOptions.containerCreateNetworkingConfig, "")
	if err != nil {
		logger.Error("Failed to create container", "error", err.Error())
		return err
	}
	logger = logger.New("container-id", respBody.ID)

	logger.Debug("Starting container", "config", *runOptions.containerStartConfig)
	if err := runner.ContainerStart(context, respBody.ID, *runOptions.containerStartConfig); err != nil {
		logger.Error("Failed to start container", "error", err.Error())
		return err
	}

	logger.Debug("Requesting container logs", "config", *runOptions.containerLogsConfig)
	out, err := runner.ContainerLogs(context, respBody.ID, *runOptions.containerLogsConfig)
	if err != nil {
		logger.Error("Failed to get container logs", "error", err.Error())
		return err
	}
	defer out.Close()
	// written, err := io.Copy(runOptions.logerWriter, out)
	written, err := stdcopy.StdCopy(runOptions.logerWriter, runOptions.logerWriter, out)
	if err != nil {
		logger.Error("Failed write container logs into logger", "error", err.Error())
		return err
	}

	logger.Debug("Waiting for container to finish")
	status, err := runner.ContainerWait(context, respBody.ID)
	if err != nil {
		logger.Error("Failed to wait for container to finish container", "error", err.Error())
		return err
	}
	logger.Debug("Container exit", "status", status)

	logger.Debug("Finished to get all logs", "bytes", written)
	return nil
}

func buildRunConfig(opt *RunOptions) (*runOptions, error) {

	writer, err := os.OpenFile(opt.LoggerFD, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	mounts := []mount.Mount{}
	for _, v := range opt.Arguments.Volumes {
		kv := strings.Split(v, ":")
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: kv[0],
			Target: kv[1],
		})
	}

	return &runOptions{
		containerCreateConfig: &container.Config{
			Image: opt.Arguments.Image,
			Cmd:   append([]string{opt.Arguments.Command}, opt.Arguments.Arguments...),
			Env:   opt.Arguments.EnvironmentVariables,
		},
		containerCreateHostConfig: &container.HostConfig{
			Mounts: mounts,
		},
		containerCreateNetworkingConfig: &network.NetworkingConfig{},
		containerStartConfig:            &types.ContainerStartOptions{},
		containerLogsConfig: &types.ContainerLogsOptions{
			Follow:     true,
			ShowStdout: true,
		},
		logerWriter: writer,
	}, nil
}
