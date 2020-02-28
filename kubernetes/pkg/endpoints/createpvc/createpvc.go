package createpvc

import (
	"context"
	"encoding/json"

	"github.com/open-integration/core/pkg/logger"
	utils "github.com/open-integration/service-catalog/kubernetes/pkg/utils"
	v1 "k8s.io/api/core/v1"
)

type (
	CreatepvcOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *CreatepvcArguments
	}
)

func Createpvc(opt CreatepvcOptions) (*CreatepvcReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	log.Debug("Building Kube client for PVC")
	client, err := utils.BuildKubeClient(*opt.Arguments.Auth.Host, *opt.Arguments.Auth.Token, *opt.Arguments.Auth.CRT, log)
	if err != nil {
		return nil, err
	}
	pvc := &v1.PersistentVolumeClaim{}
	err = json.Unmarshal([]byte(opt.Arguments.Pvc), pvc)
	if err != nil {
		return nil, err
	}
	_, err = client.CoreV1().PersistentVolumeClaims(pvc.ObjectMeta.Namespace).Create(pvc)
	if err != nil {
		return nil, err
	}
	log.Debug("PVC created")
	return &CreatepvcReturns{}, nil
}
