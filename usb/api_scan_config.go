package usb

import (
	"encoding/binary"
	"log"
)

func ScanConfigCount() {
	response := readCommand(0x02, 0x22)
	count := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Scan Configuration Count:", count)
}

func ActiveScanIndex() {
	response := readCommand(0x02, 0x23)
	index := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Active Scan Index:", index)
}

func SetActiveScanIndex(index int) {
	writeCommand(0x02, 0x024, []byte{byte(index)})
}
