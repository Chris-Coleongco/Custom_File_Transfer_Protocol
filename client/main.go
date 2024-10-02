package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
)

const packet_size = 512

// send a mp4 file to the server

func read_init_packet_from_server(conn net.Conn) {
	init_packet := make([]byte, packet_size)
	bytes_read, err := conn.Read(init_packet)
	if err != nil {
		fmt.Println()
	}

	fmt.Println(bytes_read)

	fmt.Println(init_packet)

	option_code := init_packet[0]
	response_code := init_packet[1]
	size := binary.LittleEndian.Uint64(init_packet[2:66])

	fmt.Printf("yo: %v", size)

	if option_code == 'd' && response_code == 1 {
		// successful request, server will now send file chunks
		// read logic here
		// read logic here
		// read logic here
		// read logic here

		retrieved_packet_buffer := make([]byte, packet_size)

		conn.Read(retrieved_packet_buffer)

	}
}

func construct_request(opt string, file_path string) []byte {
	request := make([]byte, packet_size)

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

	request := make([]byte, packet_size)

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

	read_init_packet_from_server(conn)

	// store the size from the init_packet in a variable, divide it by 502 (the packet size minus 10 for the header) and as you read into the packet_size buffer client side, you increment for loop by 1
}
