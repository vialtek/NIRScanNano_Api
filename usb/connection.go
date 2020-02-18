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

func Connect(vendorID uint16, productID uint16) error {
	var devices = hid.Enumerate(vendorID, productID)
	if len(devices) == 0 {
		return errors.New("No device found.")
	}

	var device, err = devices[0].Open()
	if err != nil {
		return errors.New("Device connection failed.")
	}

	connection = &Connection{
		device: device,
	}

	return nil
}

func Close() {
	log.Println("Device connection closing down.")

	if Connected() {
		connection.device.Close()
	}
}

func Connected() bool {
	return connection != nil
}

func readCommand(groupByte byte, commandByte byte) []byte {
	_, err := connection.device.Write([]byte{
		0x00,
		0xC0,
		0x00,
		0x02,
		0x00,
		commandByte,
		groupByte})

	bs := make([]byte, 64)
	_, err = connection.device.Read(bs)

	if err != nil {
		log.Println("ERROR:", err)
	}

	return bs
}
