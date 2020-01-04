package upsertVolume

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/open-integration/core/pkg/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	volumetypes "github.com/docker/docker/api/types/volume"
)

type (
	UpsertVolumeOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *UpsertVolumeArguments
	}
)

type (
	volume interface {
		VolumeCreate(context.Context, volumetypes.VolumesCreateBody) (types.Volume, error)
		VolumeList(ctx context.Context, filter filters.Args) (volumetypes.VolumesListOKBody, error)
	}
)

func UpsertVolume(opt UpsertVolumeOptions) (*UpsertVolumeReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})
	log.Debug("Connecting to daemon", "host", opt.Arguments.Host, "api-version", opt.Arguments.APIVersion)
	c, err := client.NewClient(opt.Arguments.Host, opt.Arguments.APIVersion, nil, nil)
	if err != nil {
		return nil, err
	}

	if exist := isVolumeExist(opt.Context, c, opt.Arguments.Name, filters.Args{}); exist {
		log.Debug("Volume found", "name", opt.Arguments.Name)
	} else {
		log.Debug("Volume not found, creating new volume", "name", opt.Arguments.Name)
	}

	labels := map[string]string{}
	for k, v := range opt.Arguments.Labels {
		if str, ok := v.(string); ok {
			labels[k] = str
		} else {
			log.Error("Failed to convert label value to string", "key", k, "value", v)
		}
	}
	log.Debug("Creating volume", "name", opt.Arguments.Name, "labels", opt.Arguments.Labels)
	mountPoint, err := createVolume(opt.Context, c, opt.Arguments.Name, labels)
	if err != nil {
		log.Error("Failed to create volume", "name", opt.Arguments.Name, "err", err.Error())
		return nil, err
	}
	log.Debug("Volume created", "name", opt.Arguments.Name)
	return &UpsertVolumeReturns{
		MountPoint: mountPoint,
	}, nil
}

func createVolume(context context.Context, volumeCreator volume, name string, labels map[string]string) (string, error) {
	v, err := volumeCreator.VolumeCreate(context, volumetypes.VolumesCreateBody{
		Name:   name,
		Labels: labels,
	})
	return v.Mountpoint, err
}

func isVolumeExist(context context.Context, volumeFinder volume, name string, filter filters.Args) bool {
	body, err := volumeFinder.VolumeList(context, filter)
	if err != nil {
		return false
	}
	for _, v := range body.Volumes {
		if v.Name == name {
			return true
		}
	}
	return false
}
