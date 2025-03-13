package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const confFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadConf() (fileCont Config, err error) {
	usrHome, err := getConfFilePath()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := os.ReadFile(usrHome)
	if err != nil {
		log.Fatal(err)
	}

	var confCredentials Config
	if err := json.Unmarshal(jsonData, &confCredentials); err != nil {
		log.Fatal(err)
	}

	return confCredentials, nil
}

func getConfFilePath() (path string, err error) {
	usrHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return usrHome + confFileName, nil
}

func (c *Config) SetUser(username string) (err error) {
	if username == "" {
		return fmt.Errorf("error! please enter a username")
	}
	c.CurrentUserName = username

	write(*c)

	return nil
}

func write(c Config) (err error) {
	confPath, err := getConfFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(confPath, jsonData, 0666); err != nil {
		return err
	}

	return nil
}
