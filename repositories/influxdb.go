package repositories

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/ogarraux/airmon/domain"
)

type InfluxRepo struct {
	Name string
	API  api.WriteAPIBlocking
}

// Create Repo
func NewInfluxRepo(Name string, URL string, Database string) InfluxRepo {
	c := influxdb2.NewClient(URL, "")

	api := c.WriteAPIBlocking("default", Database)

	return InfluxRepo{
		Name: Name,
		API:  api,
	}
}

func (ir InfluxRepo) StoreMeasurement(am domain.AirMeasurement) error {

	measurementPoint := influxdb2.NewPoint(
		"measurements",
		map[string]string{
			"sensor":    am.SensorID,
			"collector": am.CollectorID,
		},
		map[string]interface{}{
			"temperature_c": am.Temperature_C,
			"humidity_pct":  am.Humidity_Pct,
			"co2_ppm":       am.CO2_PPM,
			"voc_ppb":       am.VOC_PPB,
			"pm25":          am.PM25,
			"pm10":          am.PM10,
		},
		time.Now())

	var additionalValues = make(map[string]interface{})

	for _, x := range am.ExtraMeasurements {

		additionalValues[x.MeasurementName] = x.MeasurementValue
	}

	additionalMeasurementsPoint := influxdb2.NewPoint(
		"optional_measurements",
		map[string]string{
			"sensor":    am.SensorID,
			"collector": am.CollectorID,
		},
		additionalValues,
		time.Now())

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := ir.API.WritePoint(ctx, measurementPoint)
	if err != nil {
		return err
	}

	err = ir.API.WritePoint(ctx, additionalMeasurementsPoint)
	if err != nil {
		return err
	}

	ir.API.Flush(ctx)
	if err != nil {
		return err
	}
	return nil
}
