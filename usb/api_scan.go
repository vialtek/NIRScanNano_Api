package usb

import (
	"encoding/binary"
	"log"
)

func PerformScan(storeInSDCard bool) {
	var storeByte byte = 0
	if storeInSDCard {
		storeByte = 1
	}

	writeCommand(0x02, 0x18, []byte{storeByte})
}

func InterpretScan() {
	writeCommand(0x02, 0x39, []byte{})
}

func EstimatedScanTime() int {
	response := readCommand(0x02, 0x37, []byte{})
	scanTime := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Estimated Scan Time:", scanTime)
	return int(scanTime)
}

func ScanCompleted() bool {
	response := readCommand(0x02, 0x19, []byte{})
	scanCompleted := response.payload[0]

	log.Println("Scan Completed:", scanCompleted)
	return scanCompleted == 1
}

func GetFileSize() {
	response := readCommand(0x00, 0x2D, []byte{})
	fileSize := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("File Size:", fileSize)
}

func GetFileData() {
	response := readCommand(0x00, 0x2E, []byte{})
	data := response.payload

	log.Println("File Data:", data)
}
