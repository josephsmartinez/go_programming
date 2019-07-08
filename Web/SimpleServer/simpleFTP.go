package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	//T CP package can only contain up to 65495 bytes of payload.
	bufferSize int = 1024
	// 30 MIN TTL Connection Time
	connectionTime int = 1800
	// Idel TTL 3 MIN to Read
	connectionRead int = 180
	// Idel TTL 10 MIN to Write
	connectionWrite  int    = 600
	internetProtocol string = "tcp"
	port             string = "3000"
	ipAddress        string = "127.0.0.1:"
)

func userAuthentication() {

}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func sendFileToClient(conn net.Conn) {
	fmt.Println("A client has connected!")
	defer conn.Close()
	file, err := os.Open("dummyfile.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	fmt.Println("Sending filename and filesize!")
	conn.Write([]byte(fileSize))
	conn.Write([]byte(fileName))
	sendBuffer := make([]byte, bufferSize)
	fmt.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		conn.Write(sendBuffer)
	}
	fmt.Println("File has been sent, closing connection!")
	return

}

func handleConnection(conn net.Conn) {

	// Time out when connection rules reached
	deadErr := conn.SetDeadline(time.Now().Add(6000 * time.Second))
	readErr := conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	wirteErr := conn.SetWriteDeadline(time.Now().Add(60 * time.Second))
	if deadErr != nil || readErr != nil || wirteErr != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	fmt.Println("Client disconnected...")

}

func main() {

	// Listening connection params
	ln, err := net.Listen(internetProtocol, ipAddress+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server Started\nServer IP: %v%v\n", ipAddress, port)
	fmt.Println("Listening...")
	// Run server. If the connection fails, log the error, and continue listening for connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		// Handle  all connections
		fmt.Println("Cleint Connected...")
		go handleConnection(conn)
	}

	fmt.Println("Server is Stopped...")

}
