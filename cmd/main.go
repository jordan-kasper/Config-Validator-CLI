package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Define the structure of your YAML data
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func main() {
	// Read the YAML file
	data, err := os.ReadFile("testdata/testconfig.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	// Unmarshal the YAML data into the Config struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return
	}

	// Custom validation (optional)
	if config.Server.Port <= 0 || config.Server.Port > 65535 {
		fmt.Println("Invalid port number")
		return
	}

	// If no errors occurred, the YAML is valid
	fmt.Println("YAML file is valid")
	fmt.Printf("Configuration: %+v\n", config)
}
