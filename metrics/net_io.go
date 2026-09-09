package metrics

import (
	"bufio"
	"fmt"
	"strings"
)

// display graph of the bytes received and transmit
// packets recieved and transmit
// for each nic

type receiveNetStat struct {
	face       string
	bytes      uint64
	packets    uint64
	errs       uint64
	drop       uint64
	fifo       uint64
	frame      uint64
	compressed uint64
	multicast  uint64
}

type transmittedNetStat struct {
	bytes      uint64
	packets    uint64
	errs       uint64
	drop       uint64
	fifo       uint64
	colls      uint64
	carrier    uint64
	compressed uint64
}

func GetNetStats() (receiveNetStat, transmittedNetStat, error) {

	recNet := receiveNetStat{}
	transmNet := transmittedNetStat{}

	path := procPath("net", "dev")
	file, err := openPath(path)
	if err != nil {
		return receiveNetStat{}, transmittedNetStat{}, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Fields(text)

		if line[0] == "eth0:" {
			cnt, err := fmt.Sscanf(text,
				"%s %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
				&recNet.face, &recNet.bytes, &recNet.packets, &recNet.errs,
				&recNet.drop, &recNet.fifo, &recNet.frame, &recNet.compressed, &recNet.multicast,
				&transmNet.bytes, &transmNet.packets, &transmNet.errs, &transmNet.drop,
				&transmNet.fifo, &transmNet.colls, &transmNet.carrier, &transmNet.compressed,
			)
			if err != nil {
				return receiveNetStat{}, transmittedNetStat{}, fmt.Errorf("error:%v", err)
			}

			if cnt < 1 {
				return receiveNetStat{}, transmittedNetStat{}, fmt.Errorf("Could not read the file: Count < 1: %d", cnt)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return receiveNetStat{}, transmittedNetStat{}, fmt.Errorf("error:%v", err)
	}

	return recNet, transmNet, fmt.Errorf("error:%v", err)
}
