package repositories

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ogarraux/airmon/domain"
	"gopkg.in/yaml.v3"
)

type YAMLConfigMapRepo struct {
	ConfigMapFilename string
	ConfigMapYAML     ConfigMapYAML
}

type ConfigMapYAML struct {
	Interval        uint   `yaml:"interval"`
	MonitorEndpoint string `yaml:"monitor_endpoint"`
	Datastores      []struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		URL      string `yaml:"url"`
		Database string `yaml:"database"`
	}
	Sensors []struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	}
}

// Create Repo
func NewYAMLConfigMapRepo(ConfigMapFilename string) (YAMLConfigMapRepo, error) {

	file, err := os.Open(ConfigMapFilename)
	if err != nil {
		return YAMLConfigMapRepo{}, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return YAMLConfigMapRepo{}, err
	}

	configYAML := ConfigMapYAML{}

	err = yaml.Unmarshal(content, &configYAML)
	if err != nil {
		return YAMLConfigMapRepo{}, err
	}

	for _, configDS := range configYAML.Datastores {
		fmt.Println("Using datastore " + configDS.Name + " of type " + configDS.Type + " at URL " + configDS.URL + " and database " + configDS.Database)
	}

	for _, configS := range configYAML.Sensors {
		fmt.Println("Using sensor " + configS.Name + " of type " + configS.Type + " at URL " + configS.URL)
	}
	return YAMLConfigMapRepo{
		ConfigMapFilename: ConfigMapFilename,
		ConfigMapYAML:     configYAML,
	}, nil
}

func (kc YAMLConfigMapRepo) GetDatastores() []domain.DataStore {

	ds := []domain.DataStore{}

	for _, configDS := range kc.ConfigMapYAML.Datastores {
		ds = append(ds, domain.DataStore{
			Name:     configDS.Name,
			Type:     configDS.Type,
			URL:      configDS.URL,
			Database: configDS.Database,
		})
	}
	return ds
}

func (kc YAMLConfigMapRepo) GetSensors() []domain.Sensor {

	s := []domain.Sensor{}

	for _, configS := range kc.ConfigMapYAML.Sensors {
		s = append(s, domain.Sensor{
			Name: configS.Name,
			Type: configS.Type,
			URL:  configS.URL,
		})
	}

	return s
}

func (kc YAMLConfigMapRepo) GetMonitorEndpoint() string {
	return kc.ConfigMapYAML.MonitorEndpoint
}

func (kc YAMLConfigMapRepo) GetInterval() uint {
	return kc.ConfigMapYAML.Interval
}
