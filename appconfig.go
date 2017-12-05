package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	ScanInterval  int    `yaml:"scan_interval"`
	ScanTimeout   int    `yaml:"scan_timeout"`
	ScanDup       bool   `yaml:"scan_duplicate"`
	BufferSize    int    `yaml:"buffer_size"`
	MQTTAddress   string `yaml:"mqtt_address"`
	MQTTTopic     string `yaml:"mqtt_topic"`
	MQTTOnTopic   string `yaml:"mqtt_on_topic"`
	MQTTOffTopic  string `yaml:"mqtt_off_topic"`
	ThrowInterval int    `yaml:"throw_interval"`
	API_Url       string `yaml:"api_url"`
}

func (a *AppConfig) readFrom(cfg string) {

	log.Printf("Reading configuration from file %s", cfg)

	yamlFile, err := ioutil.ReadFile(cfg)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		panic(err)
	}
}

func (a AppConfig) GetScanDup() bool {
	return a.ScanDup
}
