package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 10, 10)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Panicln(err)
			}
			conn.Close()
			return
		}
		log.Printf("read %v byte from %v", n, conn.RemoteAddr().String())
		buf = buf[0:n]
		n, err = conn.Write(buf)
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("write %v byte to %v", n, conn.RemoteAddr().String())
		buf = buf[0:10]
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		log.Panicln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Panicln(err)
		}
		go handleConnection(conn)
	}

}
