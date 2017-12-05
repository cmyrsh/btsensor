package bleadapter

import (
	"context"
	"encoding/hex"
	"log"
	"time"

	"sync"

	"../bledata"
	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/examples/lib/dev"
	"github.com/pkg/errors"
	"gobot.io/x/gobot"
)

func ScanAndWait(wg sync.WaitGroup, scanInterval int, loopInterval int, dupok bool) {

	defer wg.Done()

	work := func() {

		gobot.Every(time.Duration(loopInterval)*time.Second, func() {
			// Scan for specified durantion, or until interrupted by user.
			log.Printf("Scanning for %d...\n", scanInterval)
			ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), time.Duration(scanInterval)*time.Second))
			chkErr(ble.Scan(ctx, dupok, advHandler, nil))

		})

	}

	ble_dummy := bleDummy{
		name1: "hci0",
	}

	robot := gobot.NewRobot(ble_dummy.Name(), []gobot.Connection{ble_dummy}, work)

	robot.Start()
}

func chkErr(err error) {
	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		log.Printf("done\n")
	case context.Canceled:
		log.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}

func advHandler(a ble.Advertisement) {
	if a.Connectable() {
		log.Printf("[%s] C %3d:", a.Address(), a.RSSI())
	} else {
		log.Printf("[%s] N %3d:", a.Address(), a.RSSI())
	}
	comma := ""
	if len(a.LocalName()) > 0 {
		log.Printf(" Name: %s", a.LocalName())
		comma = ","
	}
	if len(a.Services()) > 0 {
		log.Printf("%s Svcs: %v", comma, a.Services())
		comma = ","
	}
	if len(a.ManufacturerData()) > 0 {
		log.Printf("%s MD: %X", comma, a.ManufacturerData())
	}

	btd := bledata.BlueToothDevice{}

	btd.BlueToothMAC = a.Address().String()
	btd.SignalStrengthIndB = a.RSSI()
	btd.ManufacturerData = hex.EncodeToString(a.ManufacturerData())

	bledata.PushData(btd)
	log.Printf("\n")
}

type bleDummy struct {
	bd    ble.Device
	name1 string
}

func (d bleDummy) Name() (name string) {
	return d.name1
}

func (d bleDummy) SetName(name string) {
	log.Printf("SetName: %s", name)
	d.name1 = name
}

func (d bleDummy) Connect() error {
	log.Printf("Opening %s", d.Name())
	device, err := dev.NewDevice(d.Name())
	if err == nil {
		ble.SetDefaultDevice(device)
		d.bd = device
	}
	return err
}

func (d bleDummy) Finalize() error {
	if nil != d.bd {
		return d.bd.Stop()
	} else {
		log.Panicln("Device already Stopped")
		return nil
	}

}
