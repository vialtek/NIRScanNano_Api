package usb

import (
	"log"
)

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
		log.Println("[USB] readCommand error:", err)
	}

	return bs
}
