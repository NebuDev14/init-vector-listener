package listener

import (
	"bufio"
	"fmt"
	"github.com/NebuDev14/init-vector-listener/talker"
	"github.com/fatih/color"
	"log"
	"net"
	"strings"
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
	yellowWriter := color.New(color.FgYellow).FprintFunc()
	redWriter := color.New(color.FgRed).FprintFunc()

	fmt.Printf("Receiving connection from %s\n", conn.LocalAddr().String())

	defer conn.Close()

	yellowWriter(conn, "Connected to Initialization Vector Submission Platform\n")
	yellowWriter(conn, "Type in a flag string to submit..\n")

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(conn.RemoteAddr().String() + ": client disconnected")
			return
		}

		if strings.HasPrefix(message, "embsec{") {
			resTemp := make(chan *talker.Response)
			go talker.SubmitFlag(message, resTemp)
			response := <-resTemp

			if response.Msg == "Success" {
				formatColorPrint(response, conn)
			} else {
				redWriter(conn, "Invalid Flag\n")
			}

		} else {
			redWriter(conn, "Invalid Flag\n")
		}
	}

}

func formatColorPrint(res *talker.Response, conn net.Conn) {
	greenWriter := color.New(color.FgGreen).FprintFunc()
	name := color.New(color.FgGreen, color.Bold).FprintFunc()
	link := color.New(color.FgGreen, color.Underline).FprintFunc()

	greenWriter(conn, "Valid flag for ")
	name(conn, res.Name)
	greenWriter(conn, ".\nSubmission Link: ")
	link(conn, res.Link)
	greenWriter(conn, "\n")
}
