package usb

import (
	"log"
	"time"
)

type DeviceMessage struct {
	flags   byte
	seq     byte
	length  byte
	cmd     byte
	payload []byte
}

func readCommand(groupByte byte, commandByte byte) *DeviceMessage {
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
		log.Println("[USB] readCommand error:", err)
	}

	return NewDeviceMessage(bs)
}

func writeCommand(groupByte byte, commandByte byte, payload []byte) {
	header := []byte{
		0x00,
		0x40,
		0x00,
		byte(len(payload) + 2),
		0x00,
		commandByte,
		groupByte,
	}

	_, err := connection.device.Write(append(header, payload...))

	time.Sleep(10 * time.Millisecond)

	response := make([]byte, 100)
	_, err = connection.device.Read(response)

	if err != nil {
		log.Println("[USB] writeCommand error:", err)
	}
}

func NewDeviceMessage(byteStream []byte) *DeviceMessage {
	return &DeviceMessage{
		flags:   byteStream[0],
		seq:     byteStream[1],
		length:  byteStream[2],
		cmd:     byteStream[3],
		payload: byteStream[4:],
	}
}
