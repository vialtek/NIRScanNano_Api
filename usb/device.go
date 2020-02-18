package usb

import (
	"github.com/karalabe/hid"
	"log"
)

func Connect() {
	var devices = hid.Enumerate(0x0451, 0x4200)
	if len(devices) > 0 {
		var device, err = devices[0].Open()
		defer device.Close()

		if err != nil {
			log.Println("Device connection failed.", err)
		}
	} else {
		log.Println("No device found.")
	}
}
