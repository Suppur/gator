package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadConf() (fileCont Config, err error) {
	usrHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error reading file %v", err)
	}

	fmt.Printf("user home dir: %v", usrHome)

	jsonData, err := os.ReadFile(usrHome + ".gatorconfig.json")
	if err != nil {
		log.Fatal(err)
	}

	var confCredentials Config
	if err := json.Unmarshal(jsonData, &confCredentials); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DB URL: %s \n UserName: %s", confCredentials.DbURL, confCredentials.CurrentUserName)

	return confCredentials, nil
}
