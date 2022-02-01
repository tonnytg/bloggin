package database

import (
	"bloggin/pkg/logger"
	"encoding/json"
	"fmt"
	"os"
)

type DBConfg struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSLMode  string `json:"sslmode"`
}

// Config is the configuration for the database package
func Config() *DBConfg {

	var conf DBConfg
	file, err := os.Open("./config/database.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error db conf:", err)
		msg := fmt.Sprintf("error db conf: %s", err)
		logger.Msg("ERROR", msg)
	}
	return &conf
}
