package main

import (
	"fmt"
	"net"
)

// send a mp4 file to the server
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	input_buffer := make([]byte, 1024)

	fmt.Scan(&input_buffer)
	conn.Write(input_buffer)
}
