package services

import (
	"fmt"

	"github.com/ogarraux/airmon/domain"
	"github.com/ogarraux/airmon/ports"
	"github.com/ogarraux/airmon/repositories"
)

type AirMonitorService struct {
	Configuration ports.ConfigurationPort
}

func NewAirMonitorService(c ports.ConfigurationPort) AirMonitorService {
	return AirMonitorService{
		Configuration: c,
	}
}

func (ams AirMonitorService) GetConfiguration() ports.ConfigurationPort {
	return ams.Configuration
}

func (ams AirMonitorService) CollectAllMeasurements() error {
	sensors := ams.Configuration.GetSensors()
	datastores := ams.Configuration.GetDatastores()

	var measurements []domain.AirMeasurement
	for _, s := range sensors {
		if s.Type == "awair" {
			sensorCollector := repositories.NewAwairRepo(s.Name, s.URL)
			measurement, err := sensorCollector.CollectMeasurements()
			if err != nil {
				// If there's an error collecting measurement from one endpoint
				// just skip that one and continue with the rest
				fmt.Println("Error collecting measurement : " + err.Error())
				continue
			}
			measurements = append(measurements, measurement)
		}
	}
	for _, d := range datastores {
		if d.Type == "influx" {
			dataStore := repositories.NewInfluxRepo(d.Name, d.URL, d.Database)
			for _, m := range measurements {
				err := dataStore.StoreMeasurement(m)
				if err != nil {
					// If there's an error storing measurements to one destination
					// Just skip that one and continue with the rest
					fmt.Println("Error storing measurement : " + err.Error())
					continue
				}
			}
		}
	}
	return nil
}
