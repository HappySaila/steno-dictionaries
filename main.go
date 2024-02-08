package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func run() error {
	// Backup original grant_wilsn.json
	filename := "grant_wilson.json"
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	backupFolder := "backups"
	timestamp := time.Now().Format("20060102_150405")
	backupFilePath := filepath.Join(backupFolder, fmt.Sprintf("%v_%v", filename, timestamp))
	_, err = os.Stat(backupFolder)
	if os.IsNotExist(err) {
		err = os.Mkdir(backupFolder, 0755)
		if err != nil {
			return err
		}
	}
	err = os.WriteFile(backupFilePath, content, 0644)
	if err != nil {
		return err
	}

	// Unmarshal JSON data into a map, Swap keys and values and sort new keys alphabetically
	var data map[string]string
	err = json.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	swapped := make(map[string]string)
	for key, value := range data {
		swapped[value] = key
	}
	var keys []string
	for key := range swapped {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	sorted := make(map[string]string)
	for _, key := range keys {
		sorted[key] = swapped[key]
	}

	// Write sorted map back to current filename directory
	outputJSON, err := json.MarshalIndent(sorted, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, outputJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
