package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	defer connection.Close()

	uri := request(connection)
	if uri == "/open" {
		openresponse(connection)
	} else if uri == "/close" {
		closeresponse(connection)
	} else {
		response(connection)
	}
}

func request(connection net.Conn) string {
	i := 0
	var uri string
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			uri = strings.Fields(ln)[1]
		}
		if ln == "" {
			break
		}
		i++
	}
	return uri
}

func response(connection net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection, body)
}

func openresponse(connection net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Open for business</strong></body></html>`
	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection, body)
}

func closeresponse(connection net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Closing time!</strong></body></html>`
	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection, body)
}
