package main

import (
	"./bleadapter"
	"./mqttadapter"
	"./bledata"
	"flag"
	"sync"
	"log"
)

func main() {

	var config string

	flag.StringVar(&config, "config", "config.yaml", "Yaml Config for Bluetooth scanner")

	flag.Parse()

	appConfig := AppConfig{}

	appConfig.readFrom(config)

	log.Print(appConfig)

	var wg sync.WaitGroup
	wg.Add(2)

	bledata.OpenChannel(appConfig.BufferSize)

	go mqttadapter.MQTTy(appConfig.MQTTAddress, mqttadapter.Osname, appConfig.MQTTTopic, appConfig.MQTTOffTopic, wg, appConfig.ScanInterval)

	// github.com/currantlabs/ble
	//go bleadapter.ScanAndWait(wg, appConfig.ScanTimeout, appConfig.ScanInterval, appConfig.ScanDup)


	// github.com/paypal/gatt
	go bleadapter.StartScan(wg, appConfig.ScanDup)
	wg.Wait()

}
