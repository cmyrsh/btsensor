package bledata

import (
	"log"
)

var (
	diChan DataChannel
)

type BlueToothInfo struct {

	// Id of Scanning Device
	ScannerId string
	// Scan Start time in milliseconds after epoch
	ScanTimeStart int64
	// Scan End time in milliseconds after epoch
	ScanTimeEnd int64

	DeviceList []BlueToothDevice
}

type BlueToothDevice struct {

	// Bluetooth MAC Address
	BlueToothMAC string
	// Signal Strength
	SignalStrengthIndB int
	// Manufacturer Data
	ManufacturerData string
}

type DataChannel chan BlueToothDevice

func PushData(info BlueToothDevice) {
	diChan <- info
}

func OpenChannel(channelsize int) {
	diChan = make(DataChannel, channelsize)
}

func GetData() []BlueToothDevice {

	di_array := []BlueToothDevice{}
	select {
	case di := <-diChan:
		// if message found in channel then push the message to slice
		di_array = append(di_array, di)
		// collect all the messages from channel
		for len(diChan) > 0 {
			di_array = append(di_array, <-diChan)
		}


	default:
		log.Println("no device found")
	}

	return di_array
}
