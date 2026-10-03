package metrics

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type processStat struct {
	pid         int32
	processName string
	state       rune
	ppid        int32
	pgrp        int32
	session     int32
	tty_nr      int32
	tpgid       int32
	flags       uint
	minflt      uint32
	cminflt     uint64
	majflt      uint32
	cmajflt     uint64
	utime       uint64
	stime       uint64
	cutime      int32
	cstime      int64
	priority    int64
	nice        int64
	numThreads  int64
}

var processList []string

func GetProcessList() ([]string, error) {

	dirEntry, err := os.ReadDir("/proc")
	if err != nil {
		return []string{}, err
	}

	for _, v := range dirEntry {
		if matched, _ := regexp.Match(`[0-9]+`, []byte(v.Name())); matched {
			processList = append(processList, v.Name())
		}
	}

	return processList, nil
}

// returns the pointer of the processStat struct
func ShowPerProcessData() (*processStat, error) {

	procStat := &processStat{}

	for _, pid := range processList {

		processFilePath := procPath(pid, "stat")

		file, err := openPath(processFilePath)

		if err != nil {
			return &processStat{}, err
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			count, err := fmt.Sscanf(
				line,
				"%d %s %c %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
				&procStat.pid, &procStat.processName, &procStat.state, &procStat.ppid, &procStat.pgrp,
				&procStat.session, &procStat.tty_nr, &procStat.tpgid, &procStat.flags, &procStat.minflt,
				&procStat.cminflt, &procStat.majflt, &procStat.cmajflt, &procStat.utime,
				&procStat.stime, &procStat.cutime, &procStat.cstime, &procStat.priority,
				&procStat.nice, &procStat.numThreads,
			)

			if err != nil {
				return &processStat{}, fmt.Errorf("kd%v", err)
			}

			if count == 0 {
				return &processStat{}, fmt.Errorf("%v", err)
			}

			procStat.utime /= clockTick
			procStat.stime /= clockTick
			procStat.cutime /= clockTick
			procStat.cstime /= clockTick
		}

		if err := scanner.Err(); err != nil {
			return &processStat{}, err
		}
	}

	return procStat, nil
}
