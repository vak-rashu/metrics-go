package metrics

import (
	"bufio"
	"strconv"
	"strings"
)

// represents the following values only right now:
// MemTotal , MemFree, MemAvailable
type memStats struct {
	memTotal     string
	memFree      string
	MemAvailable string
}

var total int
var free int
var avail int

func GetMemStats() (int, int, int) {
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
			memInfo.memTotal = line[1]
			total, _ = strconv.Atoi(memInfo.memTotal)
			total /= 1000000
		}
		if line[0] == "MemFree:" {
			// if cnt, _ := fmt.Sscanf(text, "%f", &memInfo.memFree); cnt < 0 {
			// 	fmt.Print(memInfo.memFree)
			// 	os.Exit(1)
			// }
			memInfo.memFree = line[1]
			free, _ = strconv.Atoi(memInfo.memFree)
			free /= 1000000
		}
		if line[0] == "MemAvailable:" {
			memInfo.MemAvailable = line[1]
			avail, _ = strconv.Atoi(memInfo.MemAvailable)
			avail /= 1000000
			// if cnt, _ := fmt.Sscanf(text, "%s %s %f", &a, &b, &memInfo.MemAvailable); cnt < 0 {
			// 	os.Exit(1)
			// }
		}
	}
	scanner.Err()
	return total, free, avail
}
