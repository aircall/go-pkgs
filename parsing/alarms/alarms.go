package alarms

import (
	"fmt"
	"strings"
)

type Alarm struct {
	Name, Project, Component, Environment string
}

func Parse(a string) (Alarm, error) {
	ax := strings.Split(a, "_")

	if len(ax) != 5 {
		return Alarm{}, fmt.Errorf("Failed to parse alarm")
	}

	res := Alarm{
		Name:        a,
		Environment: ax[1],
		Project:     ax[2],
		Component:   ax[3],
	}

	if ax[3] == ax[2] {
		res.Component = "web"
	}
	return res, nil
}
