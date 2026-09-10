package metrics

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// represents the following values only right now:
// MemTotal , MemFree, MemAvailable
type memStats struct {
	memTotal     uint64
	memFree      uint64
	MemAvailable uint64
}

func GetMemStats() {
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
			if cnt, _ := fmt.Sscanf(line[1], "%d", &memInfo.memTotal); cnt < 0 {
				os.Exit(1)
			}
		}
		if line[0] == "MemFree:" {
			if cnt, _ := fmt.Sscanf(line[1], "%d", &memInfo.memFree); cnt < 0 {
				os.Exit(1)
			}
		}
		if line[0] == "MemAvailable:" {
			if cnt, _ := fmt.Sscanf(line[1], "%d", &memInfo.MemAvailable); cnt < 0 {
				os.Exit(1)
			}
		}
	}
	fmt.Println(memInfo.memFree, memInfo.memTotal, memInfo.MemAvailable)
	scanner.Err()
}
