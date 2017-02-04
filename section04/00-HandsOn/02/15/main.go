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

		serve(connection)
	}
}

func serve(connection net.Conn) {
	defer connection.Close()
	scanner := bufio.NewScanner(connection)
	var i int
	var method, uri string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			method = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("Method:", method)
			fmt.Println("URI:", uri)
		}
		if ln == "" {
			fmt.Println("This is the end of the message.")
			break
		}
		i++
	}
	body := `<html><head>
	<meta charset="UTF-8"><title>Low Level!</title></head><body><h1>Holy Cow this is low level.</h1></body></html>`

	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)

}
