package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ogarraux/airmon/domain"
)

type AwairRepo struct {
	Name string
	URL  string
}

type AwairResponse struct {
	Score           uint    `json:"score"`
	Temp            float64 `json:"temp"`
	Humid           float64 `json:"humid"`
	CO2             float64 `json:"co2"`
	CO2Est          float64 `json:"co2_est"`
	VOC             float64 `json:"voc"`
	VOC_H2_Raw      float64 `json:"voc_h2_raw"`
	VOC_Ethanol_Raw float64 `json:"voc_ethanol_raw"`
	PM25            float64 `json:"pm25"`
	PM10Est         float64 `json:"pm10_est"`
}

func NewAwairRepo(Name string, URL string) AwairRepo {
	return AwairRepo{
		Name: Name,
		URL:  URL,
	}
}

func (ar AwairRepo) CollectMeasurements() (domain.AirMeasurement, error) {

	c := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := c.Get(ar.URL)

	if err != nil {
		return AwairResponse{}.BuildAirMeasurement(), err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return AwairResponse{}.BuildAirMeasurement(), err
	}

	respObj := AwairResponse{}

	err = json.Unmarshal([]byte(body), &respObj)

	if err != nil {
		return AwairResponse{}.BuildAirMeasurement(), err
	}

	am := respObj.BuildAirMeasurement()
	am.CollectorID = "dev"
	am.SensorID = ar.Name

	return am, nil
}

func (aresp AwairResponse) BuildAirMeasurement() domain.AirMeasurement {

	scoreMeasurement := domain.AdditionalMeasurement{
		MeasurementName:  "AwairScore",
		MeasurementValue: float64(aresp.Score),
	}
	co2EstMeasurement := domain.AdditionalMeasurement{
		MeasurementName:  "CO2FromVOC",
		MeasurementValue: aresp.CO2Est,
	}
	vocH2Measurement := domain.AdditionalMeasurement{
		MeasurementName:  "VOCH2Raw",
		MeasurementValue: aresp.VOC_H2_Raw,
	}
	vocEthMeasurement := domain.AdditionalMeasurement{
		MeasurementName:  "VOCEthRaw",
		MeasurementValue: aresp.VOC_Ethanol_Raw,
	}

	am := domain.AirMeasurement{
		Timestamp:     time.Now(),
		Temperature_C: aresp.Temp,
		Humidity_Pct:  aresp.Humid,
		CO2_PPM:       aresp.CO2,
		VOC_PPB:       aresp.VOC,
		PM25:          aresp.PM25,
		PM10:          aresp.PM10Est,
		ExtraMeasurements: []domain.AdditionalMeasurement{
			scoreMeasurement, co2EstMeasurement, vocH2Measurement, vocEthMeasurement,
		},
	}

	return am
}
