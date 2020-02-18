package main

import (
	"log"

	"github.com/vialtek/NIRScanNano_Api/usb"
)

func main() {
	log.Println("NIRScanNano API started...")

	usb.Connect()
}
