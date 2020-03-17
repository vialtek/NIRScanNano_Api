package usb

import (
	"encoding/binary"
	"log"
)

func ScanConfigCount() {
	response := readCommand(0x02, 0x22, []byte{})
	count := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Scan Configuration Count:", count)
}

func ActiveScanIndex() {
	response := readCommand(0x02, 0x23, []byte{})
	index := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Active Scan Index:", index)
}

func SetActiveScanIndex(index int) {
	writeCommand(0x02, 0x24, []byte{byte(index)})
}

func EraseStoredScanConfigs() {
	writeCommand(0x02, 0x21, []byte{})
}
