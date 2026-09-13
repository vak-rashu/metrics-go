package metrics

import (
	"bufio"
	"fmt"
	"strings"
)

// represents the following values only right now:
// MemTotal , MemFree, MemAvailable

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
			// memInfo.memTotal = line[1]
			// total, _ = strconv.Atoi(memInfo.memTotal)
			// total /= 1000000
		}
		if line[0] == "MemFree:" {
			_, err := fmt.Sscanf(text, "%s %f", &name, &memInfo.memFree)
			free = convertKBtoGB(memInfo.memFree)

			if err != nil {
				return 0, 0, 0, err
			}
			// memInfo.memFree = line[1]
			// free, _ = strconv.Atoi(memInfo.memFree)
			// free /= 1000000
		}
		if line[0] == "MemAvailable:" {
			// memInfo.MemAvailable = line[1]
			// avail, _ = strconv.Atoi(memInfo.MemAvailable)
			// avail /= 1000000
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

// func GetMemStats() (int, int, int) {
// 	path := procPath("meminfo")
// 	file, err := openPath(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	memInfo := memStats{}

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		line := strings.Fields(text)

// 		if line[0] == "MemTotal:" {
// 			memInfo.memTotal = line[1]
// 			total, _ = strconv.Atoi(memInfo.memTotal)
// 			total /= 1000000
// 		}
// 		if line[0] == "MemFree:" {
// 			// if cnt, _ := fmt.Sscanf(text, "%f", &memInfo.memFree); cnt < 0 {
// 			// 	fmt.Print(memInfo.memFree)
// 			// 	os.Exit(1)
// 			// }
// 			memInfo.memFree = line[1]
// 			free, _ = strconv.Atoi(memInfo.memFree)
// 			free /= 1000000
// 		}
// 		if line[0] == "MemAvailable:" {
// 			memInfo.MemAvailable = line[1]
// 			avail, _ = strconv.Atoi(memInfo.MemAvailable)
// 			avail /= 1000000
// 			// if cnt, _ := fmt.Sscanf(text, "%s %s %f", &a, &b, &memInfo.MemAvailable); cnt < 0 {
// 			// 	os.Exit(1)
// 			// }
// 		}
// 	}
// 	scanner.Err()
// 	return total, free, avail
// }
