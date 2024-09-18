package main

import (
	"flag"
	"fmt"
	"net"
)

// send a mp4 file to the server

func construct_request(opt string, file_path string) []byte {
	request := make([]byte, 1024)

	// construct the request here with headers and payload
	return request
}

func main() {
	options := []string{"d", "u"}

	d := flag.String("d", "", "use -d for downloading from server")
	u := flag.String("u", "", "use -u for uploading to server")

	flag.Parse()

	download_path := *d
	upload_path := *u

	request := make([]byte, 1024)

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
}
