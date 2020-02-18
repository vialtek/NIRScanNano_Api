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

	device := usb.Connect(VendorID, ProductID)
	if device != nil {
		usb.TivaTemp(device)
	}
}
