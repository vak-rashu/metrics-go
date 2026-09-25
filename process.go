package metrics

import (
	"os"
	"regexp"
)

func GetProcessList() ([]string, error) {

	processList := []string{}

	dirEntry, err := os.ReadDir("/proc")
	if err != nil {
		return []string{}, err
	}

	for _, v := range dirEntry {
		if matched, _ := regexp.Match(`[0-9]+`, []byte(v.Name())); matched {
			processList = append(processList, v.Name())
		}
	}

	return processList, nil
}
