package metrics

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	filesystem "github.com/vak-rashu/metrics-go/internal"
)

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
func CountCPU() (int, int, error) {

	path := filesystem.Path("stat")
	file, err := filesystem.OpenPath(path)

	if err != nil {
		return 0, 0, fmt.Errorf("error reading the file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	cpuCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		if parts[0] == "intr" {
			break
		}

		// if parts[0] == "cpu" {
		// 	continue
		// }
		if strings.HasPrefix(parts[0], "cpu") && len(parts[0]) > 3 {
			cpuCount++
		}

	}

	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("error encountered during scanning: %v", err)
	}

	return cpuCount / 2, cpuCount, nil
}

// return system-wide CPU
func ShowCPUstat() (CPUStat, error) {

	path := filesystem.Path("stat")
	file, err := filesystem.OpenPath(path)

	if err != nil {
		return CPUStat{}, fmt.Errorf("error reading the file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	systemCPUStat := CPUStat{}

	for scanner.Scan() {
		lines := scanner.Text()
		parts := strings.Fields(lines)

		var cpu string

		if parts[0] == "cpu" {
			count, err := fmt.Sscanf(lines, "%s %f %f %f %f %f %f %f %f %f %f",
				&cpu,
				&systemCPUStat.userTime, &systemCPUStat.niceTime, &systemCPUStat.systemTime, &systemCPUStat.idleTime,
				&systemCPUStat.ioWaitTime, &systemCPUStat.irqTime, &systemCPUStat.softIRQTime, &systemCPUStat.stealTime,
				&systemCPUStat.guestTime, &systemCPUStat.guestNiceTime)
			if err != nil {
				return CPUStat{}, fmt.Errorf("Error: %v", err)
			}
			if count == 0 {
				return CPUStat{}, fmt.Errorf("Error: %v", err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return CPUStat{}, fmt.Errorf("error encountered during scanning: %v", err)
	}

	return systemCPUStat, nil
}

// return per CPU metrics
func ShowPerCpuStat() error {

	path := filesystem.Path("stat")
	file, err := filesystem.OpenPath(path)

	if err != nil {
		return fmt.Errorf("error reading the file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var cpu string
		systemCPUStat := CPUStat{}

		if len(parts) == 0 {
			continue
		}
		if parts[0] == "intr" {
			break
		}
		if strings.HasPrefix(parts[0], "cpu") {
			count, err := fmt.Sscanf(line, "%s %f %f %f %f %f %f %f %f %f %f",
				&cpu,
				&systemCPUStat.userTime, &systemCPUStat.niceTime, &systemCPUStat.systemTime, &systemCPUStat.idleTime,
				&systemCPUStat.ioWaitTime, &systemCPUStat.irqTime, &systemCPUStat.softIRQTime, &systemCPUStat.stealTime,
				&systemCPUStat.guestTime, &systemCPUStat.guestNiceTime)
			if err != nil {
				return fmt.Errorf("Error: %v", err)
			}
			if count == 0 {
				fmt.Println("stats not found")
			}

			fmt.Println(systemCPUStat)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error encountered during scanning: %v", err)
	}

	return nil
}
