package scan

import (
	"log"
	"time"

	"github.com/vialtek/NIRScanNano_Api/usb"
)

func Scan() {
	log.Println("Device scan...")

	if usb.Connected() {
		performScan()
		downloadScan()
	}
}

func performScan() {
	estimatedTime := usb.EstimatedScanTime()
	usb.PerformScan(false)

	time.Sleep(time.Duration(int(float64(estimatedTime)*0.8)) * time.Millisecond)
	for {
		if usb.ScanCompleted() {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func downloadScan() {
	usb.GetFileSize()
	usb.GetFileData()
}
