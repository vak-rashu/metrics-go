package metrics

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

type receiveNetStat struct {
	face       string
	bytes      float64
	packets    float64
	errs       float64
	drop       float64
	fifo       float64
	frame      float64
	compressed float64
	multicast  float64
}

type transmittedNetStat struct {
	bytes      float64
	packets    float64
	errs       float64
	drop       float64
	fifo       float64
	colls      float64
	carrier    float64
	compressed float64
}

func getNetStats() (receiveNetStat, transmittedNetStat, error) {

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
				"%s %f %f %f %f %f %f %f %f %f %f %f %f %f %f %f %f",
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

	return recNet, transmNet, nil
}

// get netio stats for packets received and transmitted

var packRec float64
var packTransm float64

func GetPacketsStat() (float64, float64, error) {
	if packRec == 0 && packTransm == 0 {
		oldrecPack, oldtransPack, err := getNetStats()
		if err != nil {
			panic(err)
		}

		packRec = oldrecPack.packets
		packTransm = oldtransPack.packets
	}

	time.Sleep(time.Second * 1)

	newrecPack, newtransPack, err := getNetStats()
	if err != nil {
		panic(err)
	}

	recPackDelta := newrecPack.packets - packRec
	transmPackDelta := newtransPack.packets - packTransm

	packRec = newrecPack.packets
	packTransm = newtransPack.packets

	return recPackDelta, transmPackDelta, nil
}
