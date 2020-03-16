package usb

import (
	"log"
)

func HibernateDevice() {
	writeCommand(0x03, 0x0D, []byte{})
}

func SetHibernateMode(enabled bool) {
	var enabledByte byte = 0
	if enabled {
		enabledByte = 1
	}

	writeCommand(0x03, 0x0E, []byte{enabledByte})
}

func HibernateMode() {
	response := readCommand(0x03, 0x0F, []byte{})
	hibernate := response.payload[0]

	log.Println("Hibernate Mode:", hibernate)
}

func PgaGain() {
	response := readCommand(0x02, 0x28, []byte{})
	gain := response.payload[0]

	log.Println("PGA Gain:", gain)
}
