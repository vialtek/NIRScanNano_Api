package main

import (
	"log"

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
		usb.TivaTemp()
		usb.DetectorTemp()
		usb.SoftwareVersions()
	}

	usb.Close()
}
