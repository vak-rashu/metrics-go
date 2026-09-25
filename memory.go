package metrics

import (
	"bufio"
	"fmt"
	"strings"
)

type memStats struct {
	name            string
	memTotal        float64
	memFree         float64
	memAvailable    float64
	memBuffers      float64
	memCached       float64
	memSReclaimable float64
}

func GetMemStats() (float64, float64, float64, float64, float64, error) {

	var (
		total  float64
		free   float64
		avail  float64
		cached float64
		used   float64
	)

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
			_, err := fmt.Sscanf(text, "%s %f", &memInfo.name, &memInfo.memTotal)
			total = convertKBtoGB(memInfo.memTotal)

			if err != nil {
				return 0.0, 0.0, 0.0, 0.0, 0.0, err
			}
		}
		if line[0] == "MemFree:" {
			_, err := fmt.Sscanf(text, "%s %f", &memInfo.name, &memInfo.memFree)
			free = convertKBtoGB(memInfo.memFree)

			if err != nil {
				return 0.0, 0.0, 0.0, 0.0, 0.0, err
			}
		}
		if line[0] == "MemAvailable:" {
			_, err := fmt.Sscanf(text, "%s %f", &memInfo.name, &memInfo.memAvailable)
			avail = convertKBtoGB(memInfo.memAvailable)

			if err != nil {
				return 0.0, 0.0, 0.0, 0.0, 0.0, err
			}
		}

		if line[0] == "Buffers:" {
			_, err := fmt.Sscanf(text, "%s %f", &memInfo.name, &memInfo.memBuffers)

			if err != nil {
				return 0.0, 0.0, 0.0, 0.0, 0.0, err
			}
		}

		if line[0] == "Cached:" {
			_, err := fmt.Sscanf(text, "%s %f", &memInfo.name, &memInfo.memCached)
			cached = convertKBtoGB(memInfo.memCached)

			if err != nil {
				return 0.0, 0.0, 0.0, 0.0, 0.0, err
			}
		}

		if line[0] == "SReclaimable:" {
			_, err := fmt.Sscanf(text, "%s %f", &memInfo.name, &memInfo.memSReclaimable)

			if err != nil {
				return 0.0, 0.0, 0.0, 0.0, 0.0, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0.0, 0.0, 0.0, 0.0, 0.0, err
	}

	used = convertKBtoGB(memInfo.memTotal - (memInfo.memFree + memInfo.memCached + memInfo.memBuffers + memInfo.memSReclaimable))
	return total, free, avail, cached, used, nil
}

func convertKBtoGB(val float64) float64 {
	val = val / 1048576
	return val
}

// get mem usage percentage
func MemPerc() (float64, float64, float64, float64, error) {

	total, free, avail, cached, used, err := GetMemStats()
	if err != nil {
		return 0.0, 0.0, 0.0, 0.0, err
	}
	percFree := (free / total) * 100
	percAvail := (avail / total) * 100
	percCached := (cached / total) * 100
	percUsed := (used / total) * 100

	return percFree, percAvail, percCached, percUsed, nil
}
