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
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}

		serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	var i int
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// looping through each line
		ln := scanner.Text()
		if i == 0 {
			// REQUEST
			xs := strings.Fields(ln)
			fmt.Println("METHOD:", xs[0])
			fmt.Println("URL:", xs[1])
		} else {
			// neccessary
			// otherwise stuck in eternal loop
			fmt.Println("End of request")
			break
		}

		i++
	}

	// RESPONSE
	body := "Hello world"

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
