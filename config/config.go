package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/13-bit/inky-botd/util"
)

var configDir,
	configFilePath,
	inkyImagePath,
	inkyImageScript string

type Config struct {
	BulletinBirdUrl string
}

var cfg Config

func init() {
	homeDir, _ := os.UserHomeDir()
	configDir = fmt.Sprintf("%s/.inky-botd", homeDir)

	configFilePath = fmt.Sprintf("%s/config.json", configDir)
	inkyImagePath = fmt.Sprintf("%s/botd.png", configDir)
	inkyImageScript = fmt.Sprintf("%s/inky/examples/7color/image.py", homeDir)

	if !checkConfigExists() {
		log.Printf("%s does not exist, creating...\n", configDir)
		createConfig()
	}

	loadConfig()
}

func checkConfigExists() bool {
	_, err := os.Stat(configDir)

	return !os.IsNotExist(err)
}

func createConfig() {
	if err := os.Mkdir(configDir, os.FileMode(0775)); err != nil {
		util.CheckError(err)
	}

	f, err := os.Create(configFilePath)
	util.CheckError(err)

	defer f.Close()

	cfg := Config{
		BulletinBirdUrl: "http://BULLETINBIRDURL:1313",
	}

	cfgJson, _ := json.MarshalIndent(cfg, "", "  ")

	f.Write(cfgJson)

	log.Printf("Config created. Edit BulletinBird server URL in %s...\n", configFilePath)

	os.Exit(0)
}

func loadConfig() {
	f, err := os.Open(configFilePath)
	util.CheckError(err)

	defer f.Close()

	err = json.NewDecoder(f).Decode(&cfg)
	util.CheckError(err)
}

func InkyImagePath() string {
	return inkyImagePath
}

func InkyImageScript() string {
	return inkyImageScript
}

func BulletinBirdUrl() string {
	return cfg.BulletinBirdUrl
}
