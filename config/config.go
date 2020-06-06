package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//Conf -> config.yaml structure
type Conf struct {
	CountriesHost     string `yaml:"countriesHost"`
	CountriesPath     string `yaml:"countriesPath"`
	LocaleHost        string `yaml:"localeHost"`
	LocalePath        string `yaml:"localePath"`
	IndHost           string `yaml:"indHost"`
	DistrictIndPath   string `yaml:"districtIndPath"`
	HistoricalINDPath string `yaml:"historicalINDPath"`
}

//ReadConf -> Function to read configs from yaml
func (c *Conf) ReadConf() *Conf {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
