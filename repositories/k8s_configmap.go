package repositories

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ogarraux/airmon/domain"
	"gopkg.in/yaml.v3"
)

type K8SConfigMapRepo struct {
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
func NewK8SConfigMapRepo(ConfigMapFilename string) (K8SConfigMapRepo, error) {

	file, err := os.Open(ConfigMapFilename)
	if err != nil {
		return K8SConfigMapRepo{}, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return K8SConfigMapRepo{}, err
	}

	configYAML := ConfigMapYAML{}

	err = yaml.Unmarshal(content, &configYAML)
	if err != nil {
		return K8SConfigMapRepo{}, err
	}

	for _, configDS := range configYAML.Datastores {
		fmt.Println("Using datastore " + configDS.Name + " of type " + configDS.Type + " at URL " + configDS.URL + " and database " + configDS.Database)
	}

	for _, configS := range configYAML.Sensors {
		fmt.Println("Using sensor " + configS.Name + " of type " + configS.Type + " at URL " + configS.URL)
	}
	return K8SConfigMapRepo{
		ConfigMapFilename: ConfigMapFilename,
		ConfigMapYAML:     configYAML,
	}, nil
}

func (kc K8SConfigMapRepo) GetDatastores() []domain.DataStore {

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

func (kc K8SConfigMapRepo) GetSensors() []domain.Sensor {

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

func (kc K8SConfigMapRepo) GetMonitorEndpoint() string {
	return kc.ConfigMapYAML.MonitorEndpoint
}

func (kc K8SConfigMapRepo) GetInterval() uint {
	return kc.ConfigMapYAML.Interval
}
