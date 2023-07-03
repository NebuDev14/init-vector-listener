package listener

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func start() {
	host := "localhost"
	port := 1155
	addr := fmt.Sprintf("%s:%d", host, port)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	log.Printf("Server started on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
		} else {
			go acceptClient(conn)
		}
	}

}

func acceptClient(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	
	if err != nil {
		fmt.Println(err)
	}

	conn.Close()
}