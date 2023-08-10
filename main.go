package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ogarraux/airmon/handlers"
	"github.com/ogarraux/airmon/repositories"
	"github.com/ogarraux/airmon/services"
)

func main() {

	// Uncomment this to use hardcoded static configuration
	//config, err := repositories.NewStaticConfig()

	// Comment this section out to use the static configuration above instead of the YAML file
	configMapFilename := os.Getenv("AM_CONFIG")
	if configMapFilename == "" {
		fmt.Println("AM_CONFIG env variable not set")
		os.Exit(1)
	}
	config, err := repositories.NewK8SConfigMapRepo(configMapFilename)

	if err != nil {
		fmt.Println("Error loading configuration: " + err.Error())
		os.Exit(1)
	}

	svc := services.NewAirMonitorService(config)

	runOnce := flag.Bool("once", false, "Run single collection only")
	flag.Parse()

	if *runOnce {
		// Just run service directly
		svc.CollectAllMeasurements()
	} else {
		// Do normal repeated collection scheduling
		th := handlers.NewTimerHandler(svc)
		th.Schedule()
	}

}
