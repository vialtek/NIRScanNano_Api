package usb

import (
	"encoding/binary"
	"log"
)

func EstimatedScanTime() {
	response := readCommand(0x02, 0x37)
	scanTime := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Estimated Scan Time:", scanTime)
}

func ScanCompleted() {
	response := readCommand(0x02, 0x19)
	scanCompleted := response.payload[0]

	log.Println("Scan Completed:", scanCompleted)
}
