package config

import (
	"bufio"
	"os"
	"strings"
)

func LoadConfigFile() (map[string]string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	fullPath := homeDir + "/.stockcli/config"
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
