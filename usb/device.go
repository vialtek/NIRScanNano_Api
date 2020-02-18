package usb

import (
	"log"

	"github.com/karalabe/hid"
)

func Connect(vendorID uint16, productID uint16) *hid.Device {
	var devices = hid.Enumerate(vendorID, productID)
	if len(devices) > 0 {
		var device, err = devices[0].Open()

		if err != nil {
			log.Println("Device connection failed.", err)
		}

		return device
	} else {
		log.Println("No device found.")
	}
	return nil
}

func readCommand(device *hid.Device, groupByte byte, commandByte byte) []byte {
	_, err := device.Write([]byte{
		0x00,
		0xC0,
		0x00,
		0x02,
		0x00,
		commandByte,
		groupByte})

	bs := make([]byte, 64)
	_, err = device.Read(bs)

	if err != nil {
		log.Println("ERROR:", err)
	}

	return bs
}
