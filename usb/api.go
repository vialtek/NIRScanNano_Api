package usb

import (
	"encoding/binary"
	"log"
)

func TivaTemp() {
	bs := readCommand(0x03, 0x0B)
	payload := bs[4:8]

	temperature := binary.LittleEndian.Uint16(payload)

	log.Println("Tiva Temperature:", temperature)
}
