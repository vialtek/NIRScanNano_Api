package scan

import (
	"log"

	"github.com/vialtek/NIRScanNano_Api/usb"
)

const NNO_FILE_SCAN_DATA = 0x00
const NNO_FILE_REF_CAL_COEFF = 0x02

func loadScanData() {
	getFile(NNO_FILE_SCAN_DATA)
}

func loadCalibData() {
	getFile(NNO_FILE_REF_CAL_COEFF)
}

func getFile(fileType byte) {
	log.Println("Downloading file from EVM...")
	remainingBytes := usb.GetFileSize(fileType)

	for {
		if remainingBytes <= 0 {
			break
		}

		usb.GetFileData()
		remainingBytes -= 64
	}
}
