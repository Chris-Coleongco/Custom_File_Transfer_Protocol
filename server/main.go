package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func get_file() {
	//   read the file chunk by chunk. no need to read the whole file into memory, thats stupid. just read it using a seek thingy

	// add concurrent downloads later for multiple at once

	file, err := os.Open("./server-data/data.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("err readign file")
			fmt.Println("err")
			if err.Error() == "EOF" {
				break
			}
		}

		fmt.Printf("buffer contains %d bytes: %s", n, buffer)

	}
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

func interpret_input(buffer []byte) {
	// format of incoming data: [header -> (command d or u, length of file in bytes if u), payload -> (file to download)]Accept
	//

	d := flag.String("d", "", "use -d for downloading from server")
	u := flag.String("u", "", "use -u for uploading to server")

	flag.Parse()

	download_path := *d
	upload_path := *u

	if download_path != "" {
		get_file()
	}

	if upload_path != "" {
		// get ready to receive a file
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()

	// invoke protocol on first read v
	read, err := read_conn(conn)
	if err != nil {
		fmt.Println("initial read failed")
		fmt.Println(err)
		return
	}

	interpret_input(read)

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
