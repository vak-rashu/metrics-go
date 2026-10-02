package metrics

import (
	"context"
	"fmt"

	"github.com/coreos/go-systemd/v22/dbus"
)

func GetRunServicesList() {
	ctx := context.Background()
	c, err := dbus.NewSystemConnectionContext(ctx)
	p, err := c.ListUnitsContext(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range p {
		if v.SubState == "running" {
			fmt.Println(v.Name, v.SubState)
		}
	}
}
