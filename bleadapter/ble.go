package bleadapter

import (
	"encoding/hex"
	"log"
	"sync"

	"../bledata"
	gatt "github.com/paypal/gatt"
	gattoption "github.com/paypal/gatt/examples/option"
)

var (
	scanDup = false
)

func onStateChanged(d gatt.Device, s gatt.State) {
	log.Println("State:", s)
	switch s {
	case gatt.StatePoweredOn:
		log.Println("scanning...")
		d.Scan([]gatt.UUID{}, scanDup)
		return
	default:
		d.StopScanning()
	}
}

func onPeriphDiscovered(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {

	log.Printf("\nPeripheral ID:%s, NAME:(%s), RSSI: %d\n", p.ID(), p.Name(), p.ReadRSSI())
	log.Println("  Local Name        =", a.LocalName)
	log.Println("  TX Power Level    =", a.TxPowerLevel)
	log.Println("  Manufacturer Data =", a.ManufacturerData)
	log.Println("  Service Data      =", a.ServiceData)

	// Create BluetoohDevice Object

	bldevice := bledata.BlueToothDevice{}

	bldevice.BlueToothMAC = string(p.ID())
	bldevice.SignalStrengthIndB = p.ReadRSSI()
	bldevice.ManufacturerData = hex.EncodeToString(a.ManufacturerData)

	bledata.PushData(bldevice)

}

func StartScan(wg sync.WaitGroup, dupOk bool) {

	defer wg.Done()

	scanDup = dupOk

	d, err := gatt.NewDevice(gattoption.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	// Register handlers.
	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
	select {}
}
