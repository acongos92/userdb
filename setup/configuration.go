package setup

import (
	"fmt"
	"os"
	"strconv"
)

type Configuration struct {
	DatabaseAddress  string
	DatabaseName     string
	DatabasePassword string
	DatabaseUser     string
	Port             int
}

const DEFAULT_ADDRESS = "http://localhost"
const DEFAULT_PORT = 5432
const DEFAULT_USER = "andy"
const ADDRESS_KEY = "DB_ADDRESS"
const NAME_KEY = "DB_NAME"
const PASS_KEY = "DB_PASS"
const USER_KEY = "USER"
const PORT = "PORT"

//BuildConfigurationFromEnvironment builds the configuration object, returns an error if needed config is missing
func (configuration *Configuration) BuildConfigurationFromEnvironment() error {
	address := os.Getenv(ADDRESS_KEY)
	if address == "" {
		fmt.Printf("Couldnt find env %s using default %s\n", ADDRESS_KEY, DEFAULT_ADDRESS)
		address = DEFAULT_ADDRESS
	}
	configuration.DatabaseAddress = address
	name := os.Getenv(NAME_KEY)
	if name == "" {
		return fmt.Errorf("Missing required env %s\n", NAME_KEY)
	}
	configuration.DatabaseName = name
	password := os.Getenv(PASS_KEY)
	if password == "" {
		return fmt.Errorf("Missing requried env %s\n", PASS_KEY)
	}
	port, err := strconv.ParseInt(os.Getenv(PORT), 10, 64)
	if err != nil {
		fmt.Printf("couldnt find env %s using default port %d\n", PORT, DEFAULT_PORT)
		port = DEFAULT_PORT
	}
	configuration.Port = int(port)
	user := os.Getenv(USER_KEY)
	if user == "" {
		fmt.Printf("missing env %s using default %s\n", USER_KEY, DEFAULT_USER)
		user = DEFAULT_USER
	}
	configuration.DatabaseUser = user
	configuration.DatabasePassword = password
	return nil
}
