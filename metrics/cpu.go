package metrics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type CPUStat struct {
	CPU           string
	UserTime      float64
	NiceTime      float64
	SystemTime    float64
	IdleTime      float64
	IOWaitTime    float64
	IRQTime       float64
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

	path := procPath("cpuinfo")
	file, err := openPath(path)

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

// return system-wide CPU stat
func getCPUstat() (CPUStat, error) {
	path := procPath("stat")
	file, err := openPath(path)

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
				&cpu.CPU,
				&cpu.UserTime, &cpu.NiceTime, &cpu.SystemTime, &cpu.IdleTime, &cpu.IOWaitTime,
				&cpu.IRQTime, &cpu.SoftIRQTime, &cpu.StealTime, &cpu.GuestTime, &cpu.GuestNiceTime,
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

	cpu.CPU = ""
	cpu.UserTime /= clockTick
	cpu.NiceTime /= clockTick
	cpu.SystemTime /= clockTick
	cpu.IdleTime /= clockTick
	cpu.IOWaitTime /= clockTick
	cpu.IRQTime /= clockTick
	cpu.SoftIRQTime /= clockTick
	cpu.StealTime /= clockTick
	cpu.GuestTime /= clockTick
	cpu.GuestNiceTime /= clockTick

	return cpu, nil
}

// return per CPU metrics
func getPerCpuStat() error {

	path := procPath("stat")
	file, err := openPath(path)

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
				&cpu.CPU,
				&cpu.UserTime, &cpu.NiceTime, &cpu.SystemTime, &cpu.IdleTime, &cpu.IOWaitTime,
				&cpu.IRQTime, &cpu.SoftIRQTime, &cpu.StealTime, &cpu.GuestTime, &cpu.GuestNiceTime,
			)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			if count == 0 {
				fmt.Println("stats not found")
			}
			cpu.CPU = ""
			cpu.UserTime /= clockTick
			cpu.NiceTime /= clockTick
			cpu.SystemTime /= clockTick
			cpu.IdleTime /= clockTick
			cpu.IOWaitTime /= clockTick
			cpu.IRQTime /= clockTick
			cpu.SoftIRQTime /= clockTick
			cpu.StealTime /= clockTick
			cpu.GuestTime /= clockTick
			cpu.GuestNiceTime /= clockTick

			fmt.Println(cpu)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error encountered during scanning: %v", err)
	}

	return nil
}

// initialise old cpu values with 0(nil)
// for the case when the metrics tui is started for the first time
var cpuOld []float64

var perc float64

func CalculateCPUStatForMain() (float64, []float64, []float64, error) {

	getCPUOld, err := getCPUstat()
	if err != nil {
		return 0, []float64{}, []float64{}, err
	}

	if cpuOld == nil {
		cpuOld = []float64{getCPUOld.UserTime, getCPUOld.NiceTime, getCPUOld.SystemTime, getCPUOld.IdleTime, getCPUOld.IOWaitTime,
			getCPUOld.IRQTime, getCPUOld.SoftIRQTime, getCPUOld.StealTime, getCPUOld.GuestTime, getCPUOld.GuestNiceTime}
	}

	oldCPUSum := 0.0
	for _, val := range cpuOld {
		oldCPUSum += val
	}

	// after subtracting guest and guest nice time
	oldCPUSum -= (cpuOld[8] + cpuOld[9])

	// separate out the idle and IOwait time
	oldIdleTime := cpuOld[3] + cpuOld[4]
	oldCPUTime := oldCPUSum - oldIdleTime

	getCPUNew, err := getCPUstat()
	if err != nil {
		return 0, []float64{}, []float64{}, err
	}
	// initialise the current cpu with current values
	currentCPU := []float64{getCPUNew.UserTime, getCPUNew.NiceTime, getCPUNew.SystemTime, getCPUNew.IdleTime, getCPUNew.IOWaitTime,
		getCPUNew.IRQTime, getCPUNew.SoftIRQTime, getCPUNew.StealTime, getCPUNew.GuestTime, getCPUNew.GuestNiceTime}

	// summation of all the time slices
	currentCPUSum := 0.0
	for _, val := range currentCPU {
		currentCPUSum += val
	}

	currentCPUSum -= (currentCPU[8] + currentCPU[9])
	delTotalTime := currentCPUSum - oldCPUSum

	totalIdleTime := currentCPU[3] + currentCPU[4]
	totalCPUTime := currentCPUSum - totalIdleTime

	// calculate delta values
	delCPUTime := totalCPUTime - oldCPUTime
	// delIdleTime := totalIdleTime - oldIdleTime

	// calculate utilization percentage
	perc = ((delCPUTime / delTotalTime) * 100)

	cpuOld = currentCPU
	// return perc, cpuOld, currentCPU, nil
	return perc, cpuOld, currentCPU, nil
}

func CalculateCPUStatPerc() (float64, error) {

	if cpuOld == nil {
		getCPUOld, err := getCPUstat()
		if err != nil {
			return 0, err
		}
		cpuOld = []float64{getCPUOld.UserTime, getCPUOld.NiceTime, getCPUOld.SystemTime, getCPUOld.IdleTime, getCPUOld.IOWaitTime,
			getCPUOld.IRQTime, getCPUOld.SoftIRQTime, getCPUOld.StealTime, getCPUOld.GuestTime, getCPUOld.GuestNiceTime}
	}

	time.Sleep(time.Second * 1)
	getCPUNew, err := getCPUstat()
	if err != nil {
		return 0, err
	}
	// initialise the current cpu with current values
	currentCPU := []float64{getCPUNew.UserTime, getCPUNew.NiceTime, getCPUNew.SystemTime, getCPUNew.IdleTime, getCPUNew.IOWaitTime,
		getCPUNew.IRQTime, getCPUNew.SoftIRQTime, getCPUNew.StealTime, getCPUNew.GuestTime, getCPUNew.GuestNiceTime}

	oldCPUSum := 0.0
	for _, val := range cpuOld {
		oldCPUSum += val
	}

	// after subtracting guest and guest nice time
	oldCPUSum -= (cpuOld[8] + cpuOld[9])

	// separate out the idle and IOwait time
	oldIdleTime := cpuOld[3] + cpuOld[4]
	oldCPUTime := oldCPUSum - oldIdleTime

	// summation of all the time slices
	currentCPUSum := 0.0
	for _, val := range currentCPU {
		currentCPUSum += val
	}

	currentCPUSum -= (currentCPU[8] + currentCPU[9])
	delTotalTime := currentCPUSum - oldCPUSum

	totalIdleTime := currentCPU[3] + currentCPU[4]
	totalCPUTime := currentCPUSum - totalIdleTime

	// calculate delta values
	delCPUTime := totalCPUTime - oldCPUTime

	// calculate utilization percentage
	perc = ((delCPUTime / delTotalTime) * 100)

	cpuOld = currentCPU
	return perc, nil
	// return perc, nil
}
