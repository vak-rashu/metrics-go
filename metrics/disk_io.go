package metrics

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

// disk io creates the graph of
// number of reads and writes done

// the values is taken from
// {/proc/diskstats}
// all metrics are cumulative
// except field 9

type diskStat struct {
	minor          int
	major          int
	diskName       string
	readsComp      uint64
	readsMerged    uint64
	sectorRead     uint64
	readMiliSec    uint32
	writesComp     uint64
	writesMerged   uint64
	sectorsWritten uint64
	writeMiliSec   uint32
	currIO         uint32
	miliIO         uint32
	// not adding remaining fields
}

// get the rootfs on the system
func getMnt() string {
	//make a wrapper to get the no. of mounts in the system
	cmd := exec.Command("findmnt", "-n", "-o", "SOURCE", "/")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	val := string(output)
	c := strings.Split(val, "/")
	y := strings.TrimSpace(c[2])

	return y
}

func GetDiskStats() (diskStat, error) {

	//get the struct value ready
	disk := diskStat{}
	// get mnts of the system
	mntsRoot := getMnt()

	// read proc file
	path := procPath("diskstats")
	file, err := openPath(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Fields(text)

		if line[2] == mntsRoot {
			cnt, err := fmt.Sscanf(text,
				"%d %d %s %d %d %d %d %d %d %d %d %d %d",
				&disk.minor, &disk.major, &disk.diskName,
				&disk.readsComp, &disk.readsMerged, &disk.sectorRead,
				&disk.readMiliSec, &disk.writesComp, &disk.writesMerged,
				&disk.sectorsWritten, &disk.writeMiliSec,
				&disk.currIO, &disk.miliIO,
			)
			if err != nil {
				return diskStat{}, fmt.Errorf("error:%v", err)
			}

			if cnt < 1 {
				return diskStat{}, fmt.Errorf("Could not read the file: Count < 1: %d", cnt)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return diskStat{}, fmt.Errorf("error:%v", err)
	}

	return disk, fmt.Errorf("error:%v", err)
}
