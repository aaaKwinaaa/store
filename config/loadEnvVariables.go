package config

import (
	"flag"
	"fmt"
)

var ENV *string

func LoadEnvVariables() {

	ENV = flag.String("env", "dev", "select environment")
	flag.Parse()
	fmt.Printf("This environment state \"%s\" is still running...\n", *ENV)

}
