package metrics

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/go-systemd/sdjournal"
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

// var Default []serviceStruct

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

			// default:
			// 	serviceList := serviceStruct{}
			// 	serviceList.name = v.Name
			// 	serviceList.substate = v.SubState
			// 	serviceList.activeState = v.ActiveState
			// 	serviceList.desc = v.Description
			// 	serviceList.jobID = v.JobId
			// 	serviceList.loadState = v.LoadState

			// 	Default = append(Default, serviceList)
		}
	}
}

func ReadJournalLogs(name string) {
	r, err := sdjournal.NewJournal()
	if err != nil {
		log.Fatalf("Failed to open journal: %v", err)
	}
	defer r.Close()

	// 2. Filter for a specific service (e.g., ssh.service)
	service := fmt.Sprintf("_SYSTEMD_UNIT=%s", name)
	err = r.AddMatch(service)
	if err != nil {
		log.Fatalf("Failed to add match: %v", err)
	}

	// 3. Seek to the end if you want live logs, or start from the beginning
	// For this example, we'll start from the oldest available log for this service
	r.SeekRealtimeUsec(0)

	fmt.Println("Reading logs for ssh.service:")

	for {
		c, err := r.Next()
		if err != nil {
			log.Fatalf("Error moving to next entry: %v", err)
		}
		if c == 0 {
			// No more entries available at the moment
			break
		}

		entry, err := r.GetEntry()
		if err != nil {
			continue
		}

		// Print the timestamp and the message
		timestamp := time.Unix(int64(entry.RealtimeTimestamp/1000000), 0)
		fmt.Printf("[%s] %s\n", timestamp.Format("2006-01-02 15:04:05"), entry.Fields["MESSAGE"])
	}
}
