package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const packet_size = 512

func send_init_packet(conn net.Conn, file_size int64) {
	fmt.Println("entered send_init_packet")
	fmt.Println(file_size)

	init_packet := make([]byte, packet_size)

	init_packet[0] = 'd'

	init_packet[1] = 0

	fmt.Println(uint64(file_size))

	binary.LittleEndian.PutUint64(init_packet[2:], uint64(file_size))

	fmt.Println(string(init_packet))

	conn.Write(init_packet)
}

func read_init_packet_response(conn net.Conn) int {

	buffer, err := read_conn(conn)

	if err != nil {
		fmt.Println(err)
	}

	return int(buffer[1])
}

func send_file(conn net.Conn, file *os.File, file_size int64) {

	file_reader := bufio.NewReader(file)

	var current_peak int = 511

	for {
		buffer, err := file_reader.Peek(current_peak)

		if err != nil {
			fmt.Println(err)
		}

		if int64(current_peak) < file_size {
			send_file_chunk(conn, buffer)
			current_peak += (packet_size - 1)
		} else {
			break
		}
	}

}

func send_file_chunk(conn net.Conn, file_chunk []byte) {
	conn.Write(file_chunk)
}

func get_file(conn net.Conn, requested_file string) {
	//   read the file chunk by chunk. no need to read the whole file into memory, thats stupid. just read it using a seek thingy

	// next 4 bytes will define the length of file | 32 bit integer will mean 4,294,967,295 bytes is the max size of the file to transfer, aka 4.294967295 GB

	// add concurrent downloads later for multiple at once
	fmt.Println("entered get_file")
	fmt.Println([]byte(requested_file)) // need to get out null bytes
	requested_file_trimmed := bytes.Trim([]byte(requested_file), "\x00")

	file, err := os.Open(string(requested_file_trimmed))
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	file_size := stat.Size()
	// first send a packet to the client with length of the file

	send_init_packet(conn, file_size)

	response := read_init_packet_response(conn)

	if response == 1 {
		// successful response therefore start sending file chunks

		send_file(conn, file, file_size)
	}

}

func read_conn(conn net.Conn) ([]byte, error) {
	buffer := make([]byte, packet_size)
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
	fmt.Println("entered interpret_input")
	fmt.Println(buffer)

	if string(buffer[0]) == "d" {
		// download the file and send to user
		fmt.Println(buffer)
		get_file(conn, string(buffer[1:]))
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()

	init_packet := make([]byte, packet_size)

	// invoke protocol on first read v
	read, err := conn.Read(init_packet)
	if err != nil {
		fmt.Println("initial read failed")
		fmt.Println(err)
		return
	}

	fmt.Println("AKDGHOIAEDJG")
	fmt.Println(read)

	fmt.Println(init_packet)

	interpret_input(conn, init_packet)
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
