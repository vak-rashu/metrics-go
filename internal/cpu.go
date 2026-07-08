package metrics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var statFilePath = "/proc/stat"

type CPUStat struct {
	cpu           string
	userTime      float64
	niceTime      float64
	systemTime    float64
	idleTime      float64
	ioWaitTime    float64
	irqTime       float64
	softIRQTime   float64
	stealTime     float64 // time stolen by a hypervisor
	guestTime     float64 // time spent running a virtual CPU
	guestNiceTime float64 // time spent running a niced virtual CPU
}

// return how many total logical cpus are there
func CountCPU() {
	file, err := os.Open(statFilePath)
	if err != nil {
		log.Fatalf("error reading the file:", err)
	}

	scanner := bufio.NewScanner(file)

	buf := make([]byte, 0, 8*1024)
	scanner.Buffer(buf, 1024*1024)

	n := 0
	n = n + 1

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		n = n + 1

		if parts[0] == "intr" {
			break
		}
	}

	fmt.Printf("Total number of CPU are %d\nTotal number of Logical CPU are %d\n", (n/2)-1, n-2)
}

//return system-wide CPU and per CPU stat values
// func showCPUstat{}
