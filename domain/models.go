package domain

import "time"

type AirMeasurement struct {
	Timestamp         time.Time
	SensorID          string
	CollectorID       string
	Temperature_C     float64
	Humidity_Pct      float64
	CO2_PPM           float64
	VOC_PPB           float64
	PM25              float64
	PM10              float64
	ExtraMeasurements []AdditionalMeasurement
}

type AdditionalMeasurement struct {
	MeasurementName  string
	MeasurementValue float64
}

// Configuration Related
type Sensor struct {
	Type string
	Name string
	URL  string
}

type DataStore struct {
	Type     string
	Name     string
	URL      string
	Database string
}

type Configuration struct {
	Sensors              []Sensor
	DataStores           []DataStore
	CollectionInterval_S uint
	MonitorEndpoint      string
}
