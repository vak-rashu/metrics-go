package metrics

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// get the sector size of

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

// get the block devices on the system
func GetBlockDevice() ([]string, error) {

	dirSlice := []string{}
	dirEntry, err := os.ReadDir("/sys/block")
	if err != nil {
		return []string{}, err
	}

	for _, v := range dirEntry {
		if matched, _ := regexp.Match(`sd*`, []byte(v.Name())); matched {
			dirSlice = append(dirSlice, v.Name())
		}
	}

	return dirSlice, nil
}

func getBlockSize(blockName string) (float64, error) {
	path := sysPath("block", blockName, "queue", "physical_block_size")
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	blockSize, err := strconv.Atoi(strings.Trim(string(b), "\n"))
	if err != nil {
		return 0, err
	}

	return float64(blockSize), nil
}

func getDiskStats(blockName string) (diskStat, error) {

	//get the struct value ready
	disk := diskStat{}

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

		if line[2] == blockName {
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

func GetDiskReadIOPS(blockName string) (float64, error) {

	if oldreadComp == 0 {
		olddisk, err := getDiskStats(blockName)
		if err != nil {
			return 0.0, err
		}

		oldreadComp = olddisk.readComps
		// return 0.0, nil
	}

	newDisk, err := getDiskStats(blockName)
	if err != nil {
		return 0.0, err
	}

	diskDelta := newDisk.readComps - oldreadComp

	oldreadComp = newDisk.readComps

	return diskDelta, nil
}

var oldwriteComp float64

func GetDiskWriteIOPS(blockName string) (float64, error) {

	if oldwriteComp == 0 {
		olddisk, err := getDiskStats(blockName)
		if err != nil {
			return 0.0, err
		}

		oldwriteComp = olddisk.writesComp
		// return 0.0, nil
	}

	newDisk, err := getDiskStats(blockName)
	if err != nil {
		return 0.0, err
	}

	diskDelta := newDisk.writesComp - oldwriteComp

	oldwriteComp = newDisk.writesComp

	return diskDelta, nil
}

var oldreadBytes float64

func GetDiskReadBytes(blockName string) (float64, error) {

	if oldreadBytes == 0 {
		olddisk, err := getDiskStats(blockName)
		if err != nil {
			return 0.0, err
		}

		oldreadBytes = olddisk.sectorRead
		// return 0.0, nil
	}

	newDisk, err := getDiskStats(blockName)
	if err != nil {
		return 0.0, err
	}

	b, err := getBlockSize(blockName)
	if err != nil {
		return 0.0, err
	}

	diskByteReadDelta := (newDisk.sectorRead - oldreadBytes) * b

	oldreadBytes = newDisk.readComps

	return diskByteReadDelta, nil
}

var oldwriteBytes float64

func GetDiskWriteBytes(blockName string) (float64, error) {

	if oldwriteBytes == 0 {
		olddisk, err := getDiskStats(blockName)
		if err != nil {
			return 0.0, err
		}

		oldwriteBytes = olddisk.sectorsWritten
		// return 0.0, nil
	}

	newDisk, err := getDiskStats(blockName)
	if err != nil {
		return 0.0, err
	}

	b, err := getBlockSize(blockName)
	if err != nil {
		return 0.0, err
	}

	diskByteWriteDelta := (newDisk.sectorsWritten - oldwriteBytes) * b

	oldwriteBytes = newDisk.sectorsWritten

	return diskByteWriteDelta, nil
}
