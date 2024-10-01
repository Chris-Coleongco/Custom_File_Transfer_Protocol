
package main

import (
    "fmt"
    "os"
)

func main() {
    filePath := "/home/soy/custom-protocol/server/server-data/2024-09-17 16-10-01.mkv"
    println([]byte(filePath))
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    fmt.Println(file)
    defer file.Close()

    // Proceed with your file operations
}
