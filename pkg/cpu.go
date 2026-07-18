package metrics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
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

const clockTick = 100

// return how many total logical cpus are there
func LogicalCpuCount() (int, error) {

	path := filesystem.Path("cpuinfo")
	file, err := filesystem.OpenPath(path)

	if err != nil {
		return 0, fmt.Errorf("error reading the file: %v", err)
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

		if parts[0] == "processor" {
			cpuCount++
		}

	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error encountered during scanning: %v", err)
	}

	return cpuCount, nil
}

func CountCpuCores() {

	physicalIdPath := "/sys/devices/system/cpu/cpu[0-9]*/topology/physical_package_id"

	slice1, err := filepath.Glob(physicalIdPath)

	if err != nil {
		fmt.Println(err)
		// return 0, fmt.Errorf("%v", err)
	}

	coreIdSlice := make([]int, 0)
	coresMap := make(map[int][]int)

	for _, file := range slice1 {
		sysPath := strings.Split(file, "/")
		coreIdPath := "/sys/devices/system/cpu/" + sysPath[5] + "/topology/core_id"

		physicalId, err := os.ReadFile(file)
		coreId, err := os.ReadFile(coreIdPath)

		if err != nil {
			fmt.Println(err)
		}

		intPhysicalId, err := strconv.Atoi(strings.Trim(string(physicalId), "\n"))
		intCoreId, err := strconv.Atoi(strings.Trim(string(coreId), "\n"))
		if err != nil {
			fmt.Println(err)
		}

		coreIdSlice = append(coreIdSlice, intCoreId)

		coresMap[intPhysicalId] = coreIdSlice

		slices.Sort(coresMap[intPhysicalId])
		coresMap[intPhysicalId] = slices.Compact(coresMap[intPhysicalId])
		fmt.Println(coreIdSlice, coresMap)

	}

	coresCount := 0
	for _, slice := range coresMap {
		coresCount += len(slice)
	}

	// return coresCount, nil
	fmt.Println(coresCount)
}

// return system-wide CPU
func ShowCPUstat() (CPUStat, error) {
	path := filesystem.Path("stat")
	file, err := filesystem.OpenPath(path)

	if err != nil {
		return CPUStat{}, err
	}

	defer file.Close()

	cpu := CPUStat{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if parts[0] == "cpu" {
			count, err := fmt.Sscanf(line,
				"%s %f %f %f %f %f %f %f %f %f %f",
				&cpu.cpu,
				&cpu.userTime, &cpu.systemTime, &cpu.stealTime, &cpu.softIRQTime, &cpu.niceTime,
				&cpu.irqTime, &cpu.ioWaitTime, &cpu.idleTime, &cpu.guestTime, &cpu.guestNiceTime,
			)
			if err != nil {
				return CPUStat{}, fmt.Errorf("Error: %v", err)
			}

			if count == 0 {
				return CPUStat{}, fmt.Errorf("Error: file not parsed successfully")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return CPUStat{}, fmt.Errorf("Error: %v", err)
	}

	cpu.userTime /= clockTick
	cpu.systemTime /= clockTick
	cpu.stealTime /= clockTick
	cpu.softIRQTime /= clockTick
	cpu.niceTime /= clockTick
	cpu.idleTime /= clockTick
	cpu.ioWaitTime /= clockTick
	cpu.idleTime /= clockTick
	cpu.guestNiceTime /= clockTick
	cpu.guestNiceTime /= clockTick

	return cpu, nil
}

// return per CPU metrics
func ShowPerCpuStat() error {

	path := filesystem.Path("stat")
	file, err := filesystem.OpenPath(path)

	if err != nil {
		return fmt.Errorf("error reading the file: %v", err)
	}

	defer file.Close()

	cpu := CPUStat{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}
		if parts[0] == "intr" {
			break
		}

		// loop in every cpu[num] line
		// except for the first "cpu" line
		if len(parts[0]) > 3 {
			count, err := fmt.Sscanf(line,
				"%s %f %f %f %f %f %f %f %f %f %f",
				&cpu.cpu,
				&cpu.userTime, &cpu.systemTime, &cpu.stealTime, &cpu.softIRQTime, &cpu.niceTime,
				&cpu.irqTime, &cpu.ioWaitTime, &cpu.idleTime, &cpu.guestTime, &cpu.guestNiceTime,
			)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			if count == 0 {
				fmt.Println("stats not found")
			}

			cpu.userTime /= clockTick
			cpu.systemTime /= clockTick
			cpu.stealTime /= clockTick
			cpu.softIRQTime /= clockTick
			cpu.niceTime /= clockTick
			cpu.idleTime /= clockTick
			cpu.ioWaitTime /= clockTick
			cpu.idleTime /= clockTick
			cpu.guestNiceTime /= clockTick
			cpu.guestNiceTime /= clockTick

			fmt.Println(cpu)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error encountered during scanning: %v", err)
	}

	return nil
}

//  mapping[current_info[b'physical id']] = current_info[b'cpu cores']
