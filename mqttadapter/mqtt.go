package mqttadapter

import (
	//"encoding/json"
	"log"
	"os"
	"sync"

	"../bledata"
	"encoding/json"
	"gobot.io/x/gobot"
	"time"
	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	Osname, _ = os.Hostname()
)

type mqttD struct {
	c_o MQTT.ClientOptions
	c   MQTT.Client
}

func (d mqttD) Name() (name string) {
	return d.c_o.ClientID
}

func (d mqttD) SetName(name string) {
	d.c_o.SetClientID(name)
}

func (d mqttD) Connect() error {
	if token := d.c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		return nil
	}
}

func (d mqttD) Finalize() error {
	d.c.Disconnect(10)
	return nil
}

func MQTTy(mqtt_url string, clientid string, topic string, off_topic string, wg sync.WaitGroup, secs int) {

	defer wg.Done()

	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker(mqtt_url)
	opts.SetClientID(clientid)

	opts.SetWill(off_topic, "GoingDown", 1, false)
	opts.SetCleanSession(true)

	//create and start a client using the above ClientOptions
	c1 := MQTT.NewClient(opts)

	mymqttadp := mqttD{*opts, c1}

	work := func() {

		gobot.Every(time.Duration(secs)*time.Second, func() {


			log.Println("Checking for discovered devices..")

			device_list := bledata.GetData()

			if(len(device_list) > 0) {
				bti_info := bledata.BlueToothInfo{}

				bti_info.ScannerId = Osname
				bti_info.ScanTimeStart = time.Now().UnixNano() / 10000
				bti_info.ScanTimeEnd = time.Now().UnixNano() / 10000
				bti_info.DeviceList = device_list

				publishMessage(bti_info, &mymqttadp, topic)
			}

		})

	}
	robot := gobot.NewRobot(clientid, []gobot.Connection{mymqttadp}, work)

	robot.Start()

}

func publishMessage(di bledata.BlueToothInfo, mqttAdapter *mqttD, topic string) {

	di_json, err := json.Marshal(di)
	if err != nil {
		log.Fatal(err)
	}
	token := mqttAdapter.c.Publish(topic, 1, false, di_json)

	//log.Println("waiting for token")
	if token.Wait() && token.Error() != nil {
		log.Panic(token.Error())
	} else {
		log.Printf("published: BlueTooth Data")
	}

}
