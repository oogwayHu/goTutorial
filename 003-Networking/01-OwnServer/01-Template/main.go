package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	createServer()
	// dialServer()
}

// create new server
func createServer() {
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

		handle(conn)
	}
}

func handle(conn net.Conn) {
	// read(conn)
	// write(conn)
	readAndWrite(conn)
}

// read from server
func read(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
}

// write to server
func write(conn net.Conn) {
	defer conn.Close()

	fmt.Fprintln(conn, "Hello world")
	io.WriteString(conn, "HDSLHFSDHFLKSDHKDLSFHL")
}

// read then write to server
func readAndWrite(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Fprintln(conn, "How dare you say that \t", ln)

		if ln == "" {
			fmt.Fprintln(conn, "Line is empty, quitting...")
			break
		}
	}

	fmt.Println("End of readAndWrite")
}

// -----------------------
// -----------------------

// Dial server
func dialServer() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// readDial(conn)
	writeDial(conn)
}

// Read from unrelated server
func readDial(conn net.Conn) {
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}

// Write to server
func writeDial(conn net.Conn) {
	fmt.Fprintln(conn, "I dialed you.")
}
