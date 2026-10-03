package metrics

import (
	"context"
	"fmt"

	"github.com/coreos/go-systemd/v22/dbus"
)

type serviceStruct struct {
	name        string
	substate    string
	activeState string
	desc        string
	jobID       uint32
	loadState   string
}

var Failed []serviceStruct
var Running []serviceStruct
var Dead []serviceStruct

func GetRunServicesList() {

	ctx := context.Background()
	c, err := dbus.NewSystemConnectionContext(ctx)
	p, err := c.ListUnitsContext(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range p {
		switch v.SubState {
		case "running":
			serviceList := serviceStruct{}
			serviceList.name = v.Name
			serviceList.substate = v.SubState
			serviceList.activeState = v.ActiveState
			serviceList.desc = v.Description
			serviceList.jobID = v.JobId
			serviceList.loadState = v.LoadState
			Running = append(Running, serviceList)

		case "failed":
			serviceList := serviceStruct{}
			serviceList.name = v.Name
			serviceList.substate = v.SubState
			serviceList.activeState = v.ActiveState
			serviceList.desc = v.Description
			serviceList.jobID = v.JobId
			serviceList.loadState = v.LoadState

			Failed = append(Failed, serviceList)

		case "dead":
			serviceList := serviceStruct{}
			serviceList.name = v.Name
			serviceList.substate = v.SubState
			serviceList.activeState = v.ActiveState
			serviceList.desc = v.Description
			serviceList.jobID = v.JobId
			serviceList.loadState = v.LoadState

			Dead = append(Dead, serviceList)
		}
	}
}
