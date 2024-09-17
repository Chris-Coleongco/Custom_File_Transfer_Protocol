package main

import (
	"fmt"
	"net"
)

func get_file() {
	//   read the file chunk by chunk. no need to read the whole file into memory, thats stupid. just read it using a seek thingy
	awlkergj
}

func protocol(data []byte) {
	// the header will contain length
	// parse length
	// record length in an int var to know how much data to expect

	// can reconstruct the whole by knowing the length
}

func read_conn(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return buffer, err
	}

	return buffer, nil
}

func handle_connection(conn net.Conn) {
	defer conn.Close()

	// invoke protocol on first read v
	initial_read, err := read_conn(conn)
	if err != nil {
		fmt.Println("initial read failed")
		fmt.Println(err)
		return
	}
	protocol(initial_read)
	// the initial read needs to invoke the protocol
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handle_connection(conn)

	}
}
