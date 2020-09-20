package app

import (
	"github.com/docker/docker/api/types"
	"github.com/mfigurski80/DonateCLI/api"
)

func filterContainers(list []types.Container, cond func(types.Container) bool) []types.Container {
	r := []types.Container{}
	for i := range list {
		if cond(list[i]) {
			r = append(r, list[i])
		}
	}
	return r
}

func filterJobs(list []api.JobStruct, cond func(api.JobStruct) bool) []api.JobStruct {
	r := []api.JobStruct{}
	for i := range list {
		if cond(list[i]) {
			r = append(r, list[i])
		}
	}
	return r
}
