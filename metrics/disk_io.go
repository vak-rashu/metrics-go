package metrics

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"
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
	readComps      float64
	readsMerged    float64
	sectorRead     float64
	readMiliSec    float64
	writesComp     float64
	writesMerged   float64
	sectorsWritten float64
	writeMiliSec   float64
	currIO         float64
	miliIO         float64
	// remaining fields are not added
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
				"%d %d %s %f %f %f %f %f %f %f %f %f %f",
				&disk.minor, &disk.major, &disk.diskName,
				&disk.readComps, &disk.readsMerged, &disk.sectorRead,
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

	return disk, nil
}

// get stats for disk IOPS to
// calculate how much IO is happening
// by the system per second

var oldreadComp float64

func GetDiskReadIOPS() (float64, error) {

	if oldreadComp == 0 {
		olddisk, err := GetDiskStats()
		if err != nil {
			return 0.0, err
		}

		oldreadComp = olddisk.readComps
	}

	time.Sleep(time.Second * 1)
	newDisk, err := GetDiskStats()
	if err != nil {
		return 0.0, err
	}

	diskDelta := newDisk.readComps - oldreadComp

	oldreadComp = newDisk.readComps

	return diskDelta, nil
}

var oldwriteComp float64

func GetDiskWriteIOPS() (float64, error) {

	if oldwriteComp == 0 {
		olddisk, err := GetDiskStats()
		if err != nil {
			return 0.0, err
		}

		oldwriteComp = olddisk.writesComp
	}

	time.Sleep(time.Second * 1)

	newDisk, err := GetDiskStats()
	if err != nil {
		return 0.0, err
	}

	diskDelta := newDisk.writesComp - oldwriteComp

	oldwriteComp = newDisk.writesComp

	return diskDelta, nil
}
