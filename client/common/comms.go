package common

import (
	"encoding/binary"
	"net"
)

// Sends an array of bytes by first sending its length, then the payload
func sendMessage(conn net.Conn, msg []byte) error {
	err := binary.Write(conn, binary.BigEndian, uint16(len(msg)))
	if err != nil {
		log.Criticalf("action: send_msg_length | result: fail | error: %v", err)
		return err
	}
	written := 0
	for written < len(msg) {
		n, err := conn.Write(msg[written:])
		if err != nil {
			log.Criticalf("action: send_msg | result: fail | error: %v", err)
			return err
		}
		written += n
	}
	return nil
}

// Receives an array of bytes by first receiving its length and then the message
func receiveMessage(conn net.Conn) (string, error) {
	var length uint16
	err := binary.Read(conn, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}
	buffer := make([]byte, length)
	lengthInt := int(length)
	read := 0
	for read < lengthInt {
		n, err := conn.Read(buffer[read:])
		if err != nil {
			return "", err
		}
		read += n
	}
	return string(buffer), nil
}
