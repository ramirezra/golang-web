// # Building upon the code from the previous problem:
//
// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.
//
// Add this data to your REPONSE so that this data is displayed in the browser.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	tcp, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer tcp.Close()

	for {
		connection, err := tcp.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(connection)
	}
}

func serve(connection net.Conn) {
	defer connection.Close()

	response(connection)

	response(connection)
}

func request(connection net.Conn) {
	i := 0
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			u := strings.Fields(ln)[1]
			fmt.Println(m)
			fmt.Println(u)
		}
		if ln == "" {
			fmt.Println("This is the end of the message.")
			break
		}
	}

}

func response(connection net.Conn) {
	fmt.Println("Code got here.")
	body := "Check out the response body payload."
	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/plain\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)

}
