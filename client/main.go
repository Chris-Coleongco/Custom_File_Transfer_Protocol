package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
)

// send a mp4 file to the server

func read_init_packet(conn net.Conn) {
	init_packet := make([]byte, 512)
	bytes_read, err := conn.Read(init_packet)
	if err != nil {
		fmt.Println()
	}

	fmt.Println(bytes_read)

	fmt.Println(init_packet)

	size := binary.LittleEndian.Uint64(init_packet[1:65])

	fmt.Printf("yo: %v", size)




}

func construct_request(opt string, file_path string) []byte {
	request := make([]byte, 512)

	// construct the request here with headers and payload

	switch opt {

	case "d":

		request[0] = 'd'
		copy(request[1:], []byte(file_path))

	case "u":

		// need to get file size on client's file system
		//  IMPLEMENT LATER
		//
		//  header => { opt, size, file_path }
	}
	return request
}

func main() {
	options := []string{"d", "u"}

	d := flag.String("d", "", "use -d for downloading from server")
	u := flag.String("u", "", "use -u for uploading to server")

	flag.Parse()

	download_path := *d
	upload_path := *u

	request := make([]byte, 512)

	if download_path != "" {
		request = construct_request(options[0], download_path)
	} else if upload_path != "" {
		request = construct_request(options[1], upload_path)
	} else {
		println("ultra rare error occurred")
	}

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write(request)

	// use a buffered reader here to read the init packet and the subsequent file_chunks

	read_init_packet(conn)

	// store the size from the init_packet in a variable, divide it by 502 (the packet size minus 10 for the header) and as you read into the 512 buffer client side, you increment for loop by 1
}
