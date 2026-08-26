package metrics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	filesystem "github.com/vak-rashu/metrics-go/internal"
)

type CPUStat struct {
	cpu           string
	UserTime      float64
	NiceTime      float64
	SystemTime    float64
	IdleTime      float64
	IOWaitTime    float64
	irqTime       float64
	SoftIRQTime   float64
	StealTime     float64 // time stolen by a hypervisor
	GuestTime     float64 // time spent running a virtual CPU
	GuestNiceTime float64 // time spent running a niced virtual CPU
}

// instead of assuming
// get the cpu clocktick(clock speed)
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

// returns the number of cores count
func CountCpuCores() (int, error) {

	physicalIdPath := "/sys/devices/system/cpu/cpu[0-9]*/topology/physical_package_id"
	slice1, err := filepath.Glob(physicalIdPath)

	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	coresMap := make(map[int]map[int]struct{})
	for _, file := range slice1 {

		sysPath := filepath.Dir(file)
		coreIdPath := sysPath + "/core_id"

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

		if _, ok := coresMap[intPhysicalId]; !ok {
			coresMap[intPhysicalId] = make(map[int]struct{})
		}
		coresMap[intPhysicalId][intCoreId] = struct{}{}
	}

	coresCount := 0
	for _, slice := range coresMap {
		coresCount += len(slice)
	}

	return coresCount, nil
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
				&cpu.UserTime, &cpu.SystemTime, &cpu.StealTime, &cpu.SoftIRQTime, &cpu.NiceTime,
				&cpu.irqTime, &cpu.IOWaitTime, &cpu.IdleTime, &cpu.GuestTime, &cpu.GuestNiceTime,
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

	cpu.UserTime /= clockTick
	cpu.SystemTime /= clockTick
	cpu.StealTime /= clockTick
	cpu.SoftIRQTime /= clockTick
	cpu.NiceTime /= clockTick
	cpu.IdleTime /= clockTick
	cpu.IOWaitTime /= clockTick
	cpu.IdleTime /= clockTick
	cpu.GuestNiceTime /= clockTick
	cpu.GuestNiceTime /= clockTick

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
				&cpu.UserTime, &cpu.SystemTime, &cpu.StealTime, &cpu.SoftIRQTime, &cpu.NiceTime,
				&cpu.irqTime, &cpu.IOWaitTime, &cpu.IdleTime, &cpu.GuestTime, &cpu.GuestNiceTime,
			)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			if count == 0 {
				fmt.Println("stats not found")
			}

			cpu.UserTime /= clockTick
			cpu.SystemTime /= clockTick
			cpu.StealTime /= clockTick
			cpu.SoftIRQTime /= clockTick
			cpu.NiceTime /= clockTick
			cpu.IdleTime /= clockTick
			cpu.IOWaitTime /= clockTick
			cpu.IdleTime /= clockTick
			cpu.GuestNiceTime /= clockTick
			cpu.GuestNiceTime /= clockTick

			fmt.Println(cpu)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error encountered during scanning: %v", err)
	}

	return nil
}
