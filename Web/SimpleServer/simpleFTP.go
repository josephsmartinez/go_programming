package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	//TCP package can only contain up to 65495 bytes of payload.
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

func get(conn net.Conn) {
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

func connectionParse(userCommand string) (string, error) {
	return "", errors.New("user command not found")
}

func help() string {
	options := ""
	exit := "exit - terminate ftp sessions and exit"
	cd := "cd [location] - change remote working directory"
	ls := "ls - list contents of remote directory"
	get := "get - receive file"
	help := "help - display local help information"
	xoptions := []string{exit, cd, ls, get, help}
	for _, op := range xoptions {
		options += op + " "
	}
	return options
}

func changeDir(path string) string {
	err := os.Chdir(path)
	if err != nil {
		log.Print(err)
		return err.Error()
	}
	return ""
}

func listDir() string {
	fileNames := ""
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileNames += file.Name() + " "
	}
	return fileNames
}

func listPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

func handleConnection(conn net.Conn) {

	// Time out when connection rules reached
	deadErr := conn.SetDeadline(time.Now().Add(6000 * time.Second))
	readErr := conn.SetReadDeadline(time.Now().Add(6000 * time.Second))
	wirteErr := conn.SetWriteDeadline(time.Now().Add(6000 * time.Second))
	if deadErr != nil || readErr != nil || wirteErr != nil {
		log.Fatalln("CONN TIMEOUT")
	}
	defer conn.Close()
	// Scanner will read connection type, if file type
	scanner := bufio.NewScanner(conn)

	// Client connection
	for scanner.Scan() {
		userCommand := scanner.Text()
		fmt.Printf("user entered: %v\n", userCommand)

		returnData := ""
		if strings.Compare(userCommand, "exit") == 0 {
			fmt.Fprintf(conn, "%s\n", "good bye")
			conn.Close()
		} else if strings.Compare(userCommand, "help") == 0 {
			returnData = help()
		} else if strings.Compare(strings.Split(userCommand, " ")[0], "cd") == 0 && len(strings.Split(userCommand, " ")) == 2 {
			returnData = changeDir(strings.Split(userCommand, " ")[1])
		} else if strings.Compare(userCommand, "pwd") == 0 {
			returnData = listPwd()
		} else if strings.Compare(userCommand, "ls") == 0 {
			returnData = listDir()
		} else if strings.Compare(userCommand, "get") == 0 {
			returnData = "pending command"
		} else {
			returnData = "user command not found"
		}
		fmt.Fprintf(conn, "%s\n", returnData)
	}
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
