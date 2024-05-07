package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileRelPath = "/.stockcli/config"

func LoadConfigFile() (map[string]string, error) {
	fullPath, err := getFullPath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fullPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pieces := strings.Split(line, "=")

		// Ignore malformed entries
		if len(pieces) != 2 {
			continue
		}

		config[pieces[0]] = pieces[1]
	}

	return config, nil
}

func CreateConfigFile(configMap map[string]string) error {
	fullPath, err := getFullPath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	for key, value := range configMap {
		_, err := fmt.Fprintf(file, "%s=%s\n", key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFullPath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return homeDir + fileRelPath, nil
}
