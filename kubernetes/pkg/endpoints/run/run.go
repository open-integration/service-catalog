package run

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	utils "github.com/open-integration/service-catalog/kubernetes/pkg/utils"

	"github.com/open-integration/core/pkg/logger"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
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
	log.Debug("Creating Kubernetes client")
	client, err := utils.BuildKubeClient(*opt.Arguments.Auth.Host, *opt.Arguments.Auth.Token, *opt.Arguments.Auth.CRT, log)
	if err != nil {
		return nil, err
	}
	log.Debug("Kubernetes client created")
	pod := &v1.Pod{}
	err = json.Unmarshal([]byte(opt.Arguments.Pod), pod)
	if err != nil {
		return nil, err
	}
	log.Debug("Starting pod")
	_, err = client.CoreV1().Pods(pod.ObjectMeta.Namespace).Create(pod)
	if err != nil {
		return nil, err
	}
	if opt.Arguments.Detached != nil && *opt.Arguments.Detached {
		return &RunReturns{}, nil
	}
	if err := waitForPod("Running", client, pod, log); err != nil {
		return nil, err
	}
	log.Debug("Requesting logs")
	logReq := client.CoreV1().Pods(pod.ObjectMeta.Namespace).GetLogs(pod.ObjectMeta.Name, &v1.PodLogOptions{
		Follow: true,
	})
	podLogs, err := logReq.Stream()
	if err != nil {
		log.Error("Error getting log stream")
		return nil, err
	}
	defer podLogs.Close()

	_, err = io.Copy(log.FD(), podLogs)
	client.CoreV1().Pods(pod.ObjectMeta.Namespace).Delete(pod.ObjectMeta.Name, nil)
	return &RunReturns{}, err
}

func waitForPod(status string, client *kubernetes.Clientset, pod *v1.Pod, logger logger.Logger) error {
	w, err := client.CoreV1().Pods(pod.ObjectMeta.Namespace).Watch(metav1.ListOptions{
		Watch:           true,
		ResourceVersion: pod.ResourceVersion,
		FieldSelector: fields.Set{
			"metadata.name":      pod.ObjectMeta.Name,
			"metadata.namespace": pod.ObjectMeta.Namespace,
		}.String(),
	})
	if err != nil {
		return err
	}
	for {
		select {
		case event, ok := <-w.ResultChan():
			if !ok {
				return errors.New("Failed to get event from Kubernetes")
			}
			resp := event.Object.(*v1.Pod)
			if string(resp.Status.Phase) == status {
				w.Stop()
				return nil
			}
		case <-time.After(60 * time.Second):
			w.Stop()
			return fmt.Errorf("Pod %s doesnt reached state %s for 60 secodns, exiting", pod.ObjectMeta.Name, status)
		}
	}
}
