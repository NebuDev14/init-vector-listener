package listener

import (
	"fmt"
	"log"
	"net"
	"bufio"
)

func StartListener() {
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
	fmt.Printf("Receiving connection from %s\n", conn.LocalAddr().String())
	
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": client disconnected")
			return
		}

		if message == "bad word\n" {
			fmt.Println("nooo")
			if _, err := fmt.Fprintf(conn, "that's not cool\n"); err != nil {
				fmt.Printf("Error sending back to client: %s\n", err)
			}
		}

		fmt.Print(message)
	}

}
