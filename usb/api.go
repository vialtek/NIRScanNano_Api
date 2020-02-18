package usb

import (
	"encoding/binary"
	"log"

	"github.com/karalabe/hid"
)

func TivaTemp(device *hid.Device) {
	bs := readCommand(device, 0x03, 0x0B)
	payload := bs[4:8]

	ambient := binary.LittleEndian.Uint16(payload)

	log.Println("Tiva Temperature:", ambient)
}
