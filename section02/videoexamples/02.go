package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Robinson Ramirez"
	str := fmt.Sprint(`
    <!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
    </head>
    <body>
    <h1>` + name +
		`</h1>
    </body>
    </html>
    `)
	newfile, err := os.Create("index2.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newfile.Close()

	io.Copy(newfile, strings.NewReader(str))
}
