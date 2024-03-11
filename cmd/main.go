package main

import (
	"log"

	"github.com/akshay0074700747/Project/config"
	injectdependency "github.com/akshay0074700747/Project/injectDependency"
)

func main() {

	config, err := config.LoadConfigurations()
	if err != nil {
		log.Fatal("cannot load configurations", err)
	}

	injectdependency.Initialize(config).Start(":50007")

}
