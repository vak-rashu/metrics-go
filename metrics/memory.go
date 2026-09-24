package metrics

import (
	"bufio"
	"fmt"
	"strings"
)

type memStats struct {
	memTotal     float64
	memFree      float64
	MemAvailable float64
}

var name string

var total int
var free int
var avail int

func GetMemStats() (int, int, int, error) {
	path := procPath("meminfo")
	file, err := openPath(path)
	if err != nil {
		panic(err)
	}
	memInfo := memStats{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Fields(text)

		if line[0] == "MemTotal:" {
			_, err := fmt.Sscanf(text, "%s %f", &name, &memInfo.memTotal)
			total = convertKBtoGB(memInfo.memTotal)

			if err != nil {
				return 0, 0, 0, err
			}
		}
		if line[0] == "MemFree:" {
			_, err := fmt.Sscanf(text, "%s %f", &name, &memInfo.memFree)
			free = convertKBtoGB(memInfo.memFree)

			if err != nil {
				return 0, 0, 0, err
			}
		}
		if line[0] == "MemAvailable:" {
			_, err := fmt.Sscanf(text, "%s %f", &name, &memInfo.MemAvailable)
			avail = convertKBtoGB(memInfo.MemAvailable)

			if err != nil {
				return 0, 0, 0, err
			}
		}
	}
	scanner.Err()
	return total, free, avail, nil
}

func convertKBtoGB(val float64) int {
	val = val / 1000000
	return int(val)
}
