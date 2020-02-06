//db connection info setup
package main

import (
	"fmt"
	"os"

	"github.com/acongos92/userdb/setup"
	"github.com/acongos92/userdb/users"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const BAD_CONFIG_CODE = 5

func main() {
	configuration := setup.Configuration{}
	err := configuration.BuildConfigurationFromEnvironment()
	fmt.Printf("%v", configuration)
	if err != nil {
		//todo some actual logging lol
		fmt.Println(err.Error())
		os.Exit(BAD_CONFIG_CODE)
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", configuration.DatabaseAddress, configuration.Port, configuration.DatabaseUser, configuration.DatabaseName, configuration.DatabasePassword))
	if err != nil {
		fmt.Println("Oh no couldnt connect")
		fmt.Printf("Error %s\n", err.Error())
	}
	db.AutoMigrate(&users.User{})
	db.Close()
}
