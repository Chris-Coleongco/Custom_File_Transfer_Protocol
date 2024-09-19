package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func send_init_packet(conn net.Conn, file_size int64) {
	init_packet := make([]byte, 512)

	init_packet[0] = 'd'

	binary.BigEndian.PutUint64(init_packet[1:9], uint64(file_size))

	conn.Write(init_packet)
}

func send_file_chunk(conn net.Conn, file_chunk []byte) {
	conn.Write(file_chunk)
}

func get_file(conn net.Conn, requested_file string) int64 {
	//   read the file chunk by chunk. no need to read the whole file into memory, thats stupid. just read it using a seek thingy

	// next 4 bytes will define the length of file | 32 bit integer will mean 4,294,967,295 bytes is the max size of the file to transfer, aka 4.294967295 GB

	// add concurrent downloads later for multiple at once

	file, err := os.Open(requested_file)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	file_size := stat.Size()

	// first send a packet to the client with length of the file

	send_init_packet(conn, file_size)

	//	reader := bufio.NewReader(file)

	//	file_chunk := make([]byte, 512)

	//	for {
	//		n, err := reader.Read(file_chunk)
	//		if err != nil {
	//			fmt.Println("err readign file")
	//			fmt.Println("err")
	//			if err.Error() == "EOF" {
	//				break
	//			}
	//		}
	//
	//		fmt.Printf("buffer contains %d bytes: %s", n, file_chunk)
	//
	//		send_file_chunk(conn, file_chunk)
	//	}

	return file_size
}

func read_conn(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, 512)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return buffer, err
	}

	return buffer, nil
}

func interpret_input(conn net.Conn, buffer []byte) {
	// format of incoming data: [header -> (command d or u, length of file in bytes if u), payload -> (file to download)]Accept
	//

	// first byte is opt u or d
	//

	if string(buffer[0]) == "d" {
		// download the file and send to user
		fmt.Println(buffer[1:])
		get_file(conn, string(buffer[1:]))
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

	fmt.Println(read)

	interpret_input(conn, read)

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

		fmt.Println("starting connection")
		go handle_connection(conn)

	}
}
