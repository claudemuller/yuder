package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	netaddr, err := net.ResolveIPAddr("ip4", "127.1")
	if err != nil {
		log.Fatalf("error resolving address: %v", err)
	}

	proto := "icmp"

	for {
		conn, err := net.ListenIP("ip4:"+proto, netaddr)
		if err != nil {
			log.Fatalf("error listening on address: %v", err)
		}

		buf := make([]byte, 1024)
		bytes, remoteAddr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatalf("error reading from connection: %v", err)
		}

		log.Printf("from %s: %X (%d)\n", remoteAddr.String(), buf[:bytes], len(buf))
		log.Printf("[%X][%X][%X]\n", buf[0], buf[2], buf[16:31])
	}
}

func getInt(buf []byte) string {
	var res string

	for _, b := range buf {
		res += fmt.Sprint(int(b))
	}

	return res
}
