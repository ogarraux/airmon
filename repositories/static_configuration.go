package repositories

import "github.com/ogarraux/airmon/domain"

type StaticConfigRepo struct {
}

// Create Repo
func NewStaticConfig() (StaticConfigRepo, error) {
	return StaticConfigRepo{}, nil
}

func (sc StaticConfigRepo) GetDatastores() []domain.DataStore {

	dev := domain.DataStore{
		Type:     "influx",
		Name:     "InfluxDBHost",
		URL:      "http://influxdb.host:8086",
		Database: "airmon",
	}
	ds := []domain.DataStore{
		dev,
	}
	return ds
}

func (sc StaticConfigRepo) GetSensors() []domain.Sensor {
	air1 := domain.Sensor{
		Type: "awair",
		URL:  "http://awair.device.i.p/air-data/latest",
	}
	s := []domain.Sensor{
		air1,
	}
	return s
}

func (sc StaticConfigRepo) GetMonitorEndpoint() string {
	return "https://your.uptime.host/api/push/tokenhere"
}

func (sc StaticConfigRepo) GetInterval() uint {
	return 2
}
