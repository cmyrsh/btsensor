package httpadapter

import (
	"log"
	"os"
	"sync"
	"time"

	"../bledata"
	"gobot.io/x/gobot"
	resty "gopkg.in/resty.v1"
)

var (
	Osname, _ = os.Hostname()
)

func SendData2CentralMonitoring(http_url string, wg sync.WaitGroup, interval int) {

	defer wg.Done()

	cmhandle := CentralMonitoring{
		name: http_url,
	}

	work := func() {

		gobot.Every(time.Duration(interval)*time.Second, func() {
			device_list := bledata.GetData()

			if len(device_list) > 0 {
				bti_info := bledata.BlueToothInfo{}

				bti_info.ScannerId = Osname
				bti_info.ScanTimeStart = time.Now().UnixNano() / 10000
				bti_info.ScanTimeEnd = time.Now().UnixNano() / 10000
				bti_info.DeviceList = device_list

				resp, err := resty.R().
					SetHeader("Content-Type", "application/json").
					SetHeader("Accept", "text/plain").
					SetBody(bti_info).
					Post(http_url)

				if err != nil {
					log.Printf("Error while sending %+v. Error : %+v", bti_info, err)
				}

				log.Printf("Response : %d", resp.StatusCode())

			}
		})
	}

	robot := gobot.NewRobot(cmhandle.name, []gobot.Connection{cmhandle}, work)

	robot.Start()

}

type CentralMonitoring struct {
	name string
}

func (d CentralMonitoring) Name() (name string) {
	return d.name
}

func (d CentralMonitoring) SetName(name string) {
	log.Printf("SetName: %s", name)
	d.name = name
}

func (d CentralMonitoring) Connect() error {
	log.Printf("Opening %s", d.Name())
	return nil
}

func (d CentralMonitoring) Finalize() error {
	return nil
}
