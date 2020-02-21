package usb

import (
	"log"
)

type Response struct {
	flags   byte
	seq     byte
	length  byte
	cmd     byte
	payload []byte
}

func readCommand(groupByte byte, commandByte byte) *Response {
	_, err := connection.device.Write([]byte{
		0x00,
		0xC0,
		0x00,
		0x02,
		0x00,
		commandByte,
		groupByte})

	bs := make([]byte, 64)
	_, err = connection.device.Read(bs)

	if err != nil {
		log.Println("[USB] readCommand error:", err)
	}

	return NewResponse(bs)
}

func NewResponse(byteStream []byte) *Response {
	return &Response{
		flags:   byteStream[0],
		seq:     byteStream[1],
		length:  byteStream[2],
		cmd:     byteStream[3],
		payload: byteStream[4:],
	}
}
