//db connection info setup
package main

import (
	"fmt"
	"os"

	"github.com/acongos92/userdb/setup"
)

const BAD_CONFIG_CODE = 5

func main() {
	configuration := setup.Configuration{}
	err := configuration.BuildConfigurationFromEnvironment()
	if err != nil {
		//todo some actual logging lol
		fmt.Println(err.Error())
		os.Exit(BAD_CONFIG_CODE)
	}
}
