package usb

import (
	"encoding/binary"
	"log"
)

func TivaTemp() {
	response := readCommand(0x03, 0x0B)
	temperature := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Tiva Temperature:", temperature)
}

func DetectorTemp() {
	response := readCommand(0x03, 0x00)

	ambient := binary.LittleEndian.Uint16(response.payload[0:4])
	detector := binary.LittleEndian.Uint16(response.payload[4:8])

	log.Println("Ambient Temperature:", ambient)
	log.Println("Detector Temperature:", detector)
}

func Humidity() {
	response := readCommand(0x03, 0x02)

	hdcTemp := binary.LittleEndian.Uint16(response.payload[0:4])
	humidity := binary.LittleEndian.Uint16(response.payload[4:8])

	log.Println("HDC Temperature:", hdcTemp)
	log.Println("Humidity:", humidity)
}

func BatteryVoltage() {
	response := readCommand(0x03, 0x0A)
	voltage := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Battery voltage:", voltage)
}

func Photodetector() {
	response := readCommand(0x04, 0x02)

	red := binary.LittleEndian.Uint16(response.payload[0:4])
	green := binary.LittleEndian.Uint16(response.payload[4:8])
	blue := binary.LittleEndian.Uint16(response.payload[8:12])

	log.Println("Red:", red)
	log.Println("Green:", green)
	log.Println("Blue:", blue)
}
