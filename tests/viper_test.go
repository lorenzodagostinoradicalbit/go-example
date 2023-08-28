package main

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestViperConfig(t *testing.T) {
	// Set up the test configuration file
	viper.SetConfigName("test_config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	// Set up a test configuration
	testConfig := `
	{
		"server": {
			"port": "8080"
		},
		"database": {
			"host": "localhost",
			"port": 5432,
			"user": "testuser",
			"password": "testpassword"
		}
	}
	`

	// Write the test configuration to a file
	err := os.WriteFile("test_config.json", []byte(testConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to write test configuration file: %s", err)
	}

	// Read and parse the test configuration file
	err = viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Failed to read test configuration file: %s", err)
	}

	// Access and assert the configuration values
	serverPort := viper.GetString("server.port")
	if serverPort != "8080" {
		t.Errorf("Expected server port '8080', got '%s'", serverPort)
	}

	dbHost := viper.GetString("database.host")
	if dbHost != "localhost" {
		t.Errorf("Expected database host 'localhost', got '%s'", dbHost)
	}

	dbPort := viper.GetInt("database.port")
	if dbPort != 5432 {
		t.Errorf("Expected database port '5432', got '%d'", dbPort)
	}

	dbUser := viper.GetString("database.user")
	if dbUser != "testuser" {
		t.Errorf("Expected database user 'testuser', got '%s'", dbUser)
	}

	dbPassword := viper.GetString("database.password")
	if dbPassword != "testpassword" {
		t.Errorf("Expected database password 'testpassword', got '%s'", dbPassword)
	}

	// Clean up the test configuration file
	err = os.Remove("test_config.json")
	if err != nil {
		t.Fatalf("Failed to remove test configuration file: %s", err)
	}
}
