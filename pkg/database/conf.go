package database

import (
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
func Config() *DBConfg{
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := DBConfg{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error db conf:", err)
	}

	return &conf
}
