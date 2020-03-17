package main

import (
	"log"

	"github.com/vialtek/NIRScanNano_Api/scan"
	"github.com/vialtek/NIRScanNano_Api/usb"
)

const (
	VendorID  = 0x0451
	ProductID = 0x4200
)

func main() {
	log.Println("NIRScanNano API started...")

	var err = usb.Connect(VendorID, ProductID)
	if err != nil {
		log.Fatal(err)
	}

	if usb.Connected() {
		usb.ScanConfigCount()
		usb.SetActiveScanIndex(0)
		usb.ActiveScanIndex()
		scan.Scan()
	}

	usb.Close()
}
