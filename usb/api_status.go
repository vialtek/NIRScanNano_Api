package usb

import (
	"encoding/binary"
	"fmt"
	"log"
)

func DeviceStatus() {
	response := readCommand(0x04, 0x03, []byte{})
	status := binary.LittleEndian.Uint16(response.payload[0:4])

	log.Println("Device status:", status)
}

func SoftwareVersions() {
	response := readCommand(0x02, 0x16, []byte{})

	tivaVersion := fmt.Sprintf("%v.%v.%v", response.payload[2], response.payload[1], response.payload[0])
	dlpcSWVersion := fmt.Sprintf("%v.%v.%v", response.payload[7], response.payload[6], response.payload[5])
	dlpcFlashVersion := fmt.Sprintf("%v.%v.%v", response.payload[11], response.payload[10], response.payload[9])
	specLibVersion := fmt.Sprintf("%v.%v.%v", response.payload[15], response.payload[14], response.payload[13])
	calDataVersion := fmt.Sprintf("%v.%v.%v", response.payload[19], response.payload[18], response.payload[17])
	refCalDataVersion := fmt.Sprintf("%v.%v.%v", response.payload[23], response.payload[22], response.payload[21])
	cfgDataVersion := fmt.Sprintf("%v.%v.%v", response.payload[27], response.payload[26], response.payload[25])

	log.Println("Tiva version:", tivaVersion)
	log.Println("DLPC version:", dlpcSWVersion)
	log.Println("DLPC flash version:", dlpcFlashVersion)
	log.Println("Spectrum library version:", specLibVersion)
	log.Println("Calibration data version:", calDataVersion)
	log.Println("Ref calibration data version:", refCalDataVersion)
	log.Println("CFG data version:", cfgDataVersion)
}
