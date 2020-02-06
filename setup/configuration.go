package setup

import (
	"fmt"
	"os"
)

type Configuration struct {
	DatabaseAddress  string
	DatabaseName     string
	DatabasePassword string
}

const DEFAULT_ADDRESS = "http://localhost"
const ADDRESS_KEY = "DB_ADDRESS"
const NAME_KEY = "DB_NAME"
const PASS_KEY = "DB_PASS"

//BuildConfigurationFromEnvironment builds the configuration object, returns an error if needed config is missing
func (configuration *Configuration) BuildConfigurationFromEnvironment() error {
	address := os.Getenv(ADDRESS_KEY)
	if address == "" {
		address = DEFAULT_ADDRESS
	}
	configuration.DatabaseAddress = address
	name := os.Getenv(NAME_KEY)
	if name == "" {
		return fmt.Errorf("Missing required env database name")
	}
	configuration.DatabaseName = name
	password := os.Getenv(PASS_KEY)
	if password == "" {
		return fmt.Errorf("Missing requried env password")
	}
	configuration.DatabasePassword = password
	return nil
}
