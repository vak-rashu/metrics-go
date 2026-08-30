package metrics

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/vak-rashu/levenshtein-cli"
)

// metrics show process <process-name>

type processStat struct {
	pid        int32
	comm       string
	state      rune
	ppid       int32
	pgrp       int32
	session    int32
	tty_nr     int32
	tpgid      int32
	flags      uint
	minflt     uint32
	cminflt    uint64
	majflt     uint32
	cmajflt    uint64
	utime      uint64
	stime      uint64
	cutime     int32
	cstime     int64
	priority   int64
	nice       int64
	numThreads int64
}

// returns the process stats
func ShowPerProcessData(arg string) ([]string, processStat, error) {
	// _, pid := matchProcess(arg)

	list, pid := levenshtein.Loop(arg)
	if pid != 0 {
		pidString := strconv.Itoa(pid)

		processFilePath := procPath(fmt.Sprintf("/%s/stat", pidString))
		file, err := openPath(processFilePath)
		if err != nil {
			return nil, processStat{}, err
		}

		defer file.Close()

		procStat := processStat{}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			count, err := fmt.Sscanf(
				line,
				"%d %s %c %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
				&procStat.pid, &procStat.comm, &procStat.state, &procStat.ppid, &procStat.pgrp,
				&procStat.session, &procStat.tty_nr, &procStat.tpgid, &procStat.flags, &procStat.minflt,
				&procStat.cminflt, &procStat.majflt, &procStat.cmajflt, &procStat.utime,
				&procStat.stime, &procStat.cutime, &procStat.cstime, &procStat.priority,
				&procStat.nice, &procStat.numThreads,
			)

			if err != nil {
				return nil, processStat{}, fmt.Errorf("kd%v", err)
			}

			if count == 0 {
				return nil, processStat{}, fmt.Errorf("%v", err)
			}

			procStat.utime /= clockTick
			procStat.stime /= clockTick
			procStat.cutime /= clockTick
			procStat.cstime /= clockTick
		}

		if err := scanner.Err(); err != nil {
			return nil, processStat{}, err
		}

		return nil, procStat, nil

	} else {
		return list, processStat{}, nil
	}
}
