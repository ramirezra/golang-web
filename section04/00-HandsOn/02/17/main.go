// # Building upon the code from the previous problem:
//
// Add code to respond to the following METHODS & ROUTES:
// 	GET /
// 	GET /apply
// 	POST /apply

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
	var body string
	if method == "GET" && uri == "/" {
		body = `<html><head><meta charset="UTF-8"><title>Index</title></head><body><h1>Holy Cow this is low level.</h1></body></html>`

		io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(connection, "Content-Type: text/html\r\n")
		io.WriteString(connection, "\r\n")
		io.WriteString(connection, body)
	}
	if method == "GET" && uri == "/apply" {
		body = `<html><head><meta charset="UTF-8"><title>Apply</title></head><body>
		<form action="/apply" method="post">
			<input type="text" name="fname" placeholder="first name" autofocus autocomplete="off">
			<input type="submit" name="submit" value="click to submit">
		</form>
		</body></html>`

		io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(connection, "Content-Type: text/html\r\n")
		io.WriteString(connection, "\r\n")
		io.WriteString(connection, body)
	}
	if method == "POST" && uri == "/apply" {
		body = `<html><head><meta charset="UTF-8"><title>Post</title></head><body><h1>Form submitted!</h1></body></html>`

		io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(connection, "Content-Type: text/html\r\n")
		io.WriteString(connection, "\r\n")
		io.WriteString(connection, body)
	}
}
