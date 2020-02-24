package run

import (
	"context"
	"encoding/json"

	"github.com/open-integration/core/pkg/logger"
	"github.com/open-integration/core/pkg/utils"
	v1 "k8s.io/api/core/v1"
)

type (
	RunOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *RunArguments
	}
)

func Run(opt RunOptions) (*RunReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	kube := &utils.Kubernetes{}
	client, err := kube.BuildClient("")
	if err != nil {
		return nil, err
	}
	pod := &v1.Pod{}
	err = json.Unmarshal([]byte(*opt.Arguments.Pod), pod)
	if err != nil {
		return nil, err
	}
	log.Debug("Starting pod")
	_, err = client.CoreV1().Pods(pod.ObjectMeta.Namespace).Create(pod)
	if err != nil {
		return nil, err
	}
	return &RunReturns{}, nil
}
