package kubernetes

import (
	b64 "encoding/base64"

	"github.com/open-integration/core/pkg/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func BuildKubeClient(host string, token string, b64crt string, log logger.Logger) (*kubernetes.Clientset, error) {
	ca, err := b64.StdEncoding.DecodeString(b64crt)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(&rest.Config{
		Host:        host,
		BearerToken: token,
		TLSClientConfig: rest.TLSClientConfig{
			CAData: ca,
		},
	})
}
