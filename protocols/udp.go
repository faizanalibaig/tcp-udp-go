package protocols

import (
	"log"
	"net"
)

func udp() {
	conn, err := net.ListenPacket("udp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer func(packetConn net.PacketConn) {
		err := packetConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	buf := make([]byte, 1024)

	for {
		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received %s from %s", string(buf), addr)
		conn.WriteTo([]byte("Hello from UDP server"), addr)
	}
}
