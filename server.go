package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"syscall"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 10, 10)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF && err != syscall.ECONNRESET {
				//log.Panicln(err)
				log.Println(err)
			}
			conn.Close()
			return
		}
		log.Printf("read %v byte from %v", n, conn.RemoteAddr().String())
		buf = buf[0:n]
		n, err = conn.Write(buf)
		if err != nil {
			if !(err == syscall.EPIPE || err == syscall.ECONNRESET) {
				//log.Panicln(err)
				log.Println(err)
			}
		}
		log.Printf("write %v byte to %v", n, conn.RemoteAddr().String())
		buf = buf[0:10]
	}
}

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		log.Panicln(err)
	}
	fmt.Println("Listen on 8080....")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Panicln(err)
		}
		go handleConnection(conn)
	}

}
