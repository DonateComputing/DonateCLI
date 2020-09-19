package app

import "github.com/docker/docker/api/types"

func filter(containers []types.Container, cond func(types.Container) bool) []types.Container {
	r := []types.Container{}
	for i := range containers {
		if cond(containers[i]) {
			r = append(r, containers[i])
		}
	}
	return r
}
