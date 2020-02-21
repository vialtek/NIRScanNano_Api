package usb

import (
	"encoding/binary"
	"fmt"
	"log"
)

func TivaTemp() {
	bs := readCommand(0x03, 0x0B)
	payload := bs[4:8]

	temperature := binary.LittleEndian.Uint16(payload)

	log.Println("Tiva Temperature:", temperature)
}

func DetectorTemp() {
	bs := readCommand(0x03, 0x00)

	ambient := binary.LittleEndian.Uint16(bs[4:8])
	detector := binary.LittleEndian.Uint16(bs[8:12])

	log.Println("Ambient Temperature:", ambient)
	log.Println("Detector Temperature:", detector)
}

func SoftwareVersions() {
	bs := readCommand(0x02, 0x16)
	payload := bs[4:]

	tivaVersion := fmt.Sprintf("%v.%v.%v", payload[2], payload[1], payload[0])
	dlpcSWVersion := fmt.Sprintf("%v.%v.%v", payload[7], payload[6], payload[5])
	dlpcFlashVersion := fmt.Sprintf("%v.%v.%v", payload[11], payload[10], payload[9])
	specLibVersion := fmt.Sprintf("%v.%v.%v", payload[15], payload[14], payload[13])
	calDataVersion := fmt.Sprintf("%v.%v.%v", payload[19], payload[18], payload[17])
	refCalDataVersion := fmt.Sprintf("%v.%v.%v", payload[23], payload[22], payload[21])
	cfgDataVersion := fmt.Sprintf("%v.%v.%v", payload[27], payload[26], payload[25])

	log.Println("Tiva version:", tivaVersion)
	log.Println("DLPC version:", dlpcSWVersion)
	log.Println("DLPC flash version:", dlpcFlashVersion)
	log.Println("Spectrum library versuin:", specLibVersion)
	log.Println("Calibration data version:", calDataVersion)
	log.Println("Ref calibration data version:", refCalDataVersion)
	log.Println("CFG data version:", cfgDataVersion)
}
