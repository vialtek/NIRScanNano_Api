package usb

import (
	"errors"
	"log"

	"github.com/karalabe/hid"
)

type Connection struct {
	device *hid.Device
}

var connection *Connection

func Connected() bool {
	return connection != nil
}

func Connect(vendorID uint16, productID uint16) error {
	var devices = hid.Enumerate(vendorID, productID)
	if len(devices) == 0 {
		return errors.New("[USB] No device found.")
	}

	var device, err = devices[0].Open()
	if err != nil {
		return errors.New("[USB] Device connection failed.")
	}

	connection = &Connection{
		device: device,
	}

	return nil
}

func Close() {
	log.Println("[USB] Device connection closing down.")

	if Connected() {
		connection.device.Close()
	}
}
