package socketcomm

import (
	"fmt"
	"log"
	"net"
)

func SocketClient() {
	fmt.Println("I am socket client")
	netAddress := "localhost:9002"
	networkProtocol := "tcp"
	conn, err := net.Dial(networkProtocol, netAddress)
	if err != nil {
		log.Printf("unable to connect network at %s, error -%s", netAddress, err.Error())
		return
	}

	defer conn.Close()

	_, err = conn.Write([]byte("Hello server I am socket client"))
	if err != nil {
		log.Printf("unable to write connection, error -%s", err.Error())
	}
	readBuff := make([]byte, 1024)
	readDataLen, err := conn.Read(readBuff)
	if err != nil {
		log.Printf("unable to read connection, error -%s", err.Error())
	}
	fmt.Printf("recived: %s", string(readBuff[:readDataLen]))

}
