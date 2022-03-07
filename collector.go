//todo сделать так чтоб это работало

package main

import (
	"encoding/json"
	cp "github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const configName = "collector.json"

type Collector struct {
	DotFilesConfigDir                  string   `json:"dot-files-config-dir"`
	ForWhichUtilitiesConfigsNeedToTake []string `json:"for-which-utilities-configs-need-to-take"`
	UserDirectory                      string   `json:"user-directory"`
	FromUserDirectory                  []string `json:"from-user-directory"`
	WhereToSaveBackupConfigs           string   `json:"where-to-save-backup-configs"`
	OverwriteFiles                     bool     `json:"overwrite-files"`
	ArchiveBackupFiles                 bool     `json:"archive-backup-files"`
}

func generateDefaultConfig() *Collector {
	return &Collector{
		DotFilesConfigDir:                  "~/.config/",
		ForWhichUtilitiesConfigsNeedToTake: []string{},
		UserDirectory:                      "~/",
		FromUserDirectory:                  []string{},
		WhereToSaveBackupConfigs:           "./dotfiles/",
		OverwriteFiles:                     true,
		ArchiveBackupFiles:                 true,
	}
}

// createConfigFile creates a config file in the current directory
func createConfigFile(config *Collector) error {
	configFile, err := os.Create(configName)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	str, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(configName, str, 0600); err != nil {
		return err
	}

	return nil
}

func NewCollector() *Collector {
	JSONFile, err := ioutil.ReadFile(configName)
	if err != nil {
		err := createConfigFile(generateDefaultConfig())
		if err != nil {
			panic(err)
		}

		println("Configuration file created collector.json", "Please, edit it and restart the program")
		os.Exit(1)
	}

	var collectorConfig Collector
	if err = json.Unmarshal(JSONFile, &collectorConfig); err != nil {
		panic(err)
	}

	return &collectorConfig
}

// Backup - backup configs
func (c *Collector) Backup() {
	// create dir
	if _, err := os.Stat(c.WhereToSaveBackupConfigs); os.IsNotExist(err) {
		err := os.MkdirAll(c.WhereToSaveBackupConfigs, 0700)
		if err != nil {
			panic(err)
		}
	}

	//src := c.DotFilesConfigDir + "{" + Implode(",", c.ForWhichUtilitiesConfigsNeedToTake) + "}"

	for _, config := range c.ForWhichUtilitiesConfigsNeedToTake {
		dest := c.WhereToSaveBackupConfigs + config
		src := c.DotFilesConfigDir + config
		if _, err := os.Stat(dest); os.IsNotExist(err) {
			err := os.MkdirAll(dest, 0700)
			if err != nil {
				panic(err)
			}
		}

		cp.Copy(src, dest)
		println("Backing up", src, "to", c.WhereToSaveBackupConfigs+config)
	}

	//src = c.UserDirectory + "{" + Implode(",", c.FromUserDirectory) + "}"
	//
	//if err := cmd.Run(); err != nil {
	//	output, _ := cmd.Output()
	//	panic(string(output))
	//}
}

func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

func main() {
	c := NewCollector()
	c.Backup()
}
