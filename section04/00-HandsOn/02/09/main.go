// # Building upon the code from the previous problem:
//
// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".
//
// Pass the connection of type net.Conn as an argument into this function.
//
// Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	tcp, err := net.Listen("tcp", ":8080")
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

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			fmt.Println("This is the end of the message.")
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(connection, "I see you connected.")

}
