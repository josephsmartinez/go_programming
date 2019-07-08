package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var userName = "joseph"

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalln(err)
		log.Println("Check connection settings")
		os.Exit(0)
	}
	defer conn.Close()

	// bs, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(bs))

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Connection Successful...\n")
	for {
		// Send Request
		fmt.Print(userName + "client# ")
		userCommands, _ := reader.ReadString('\n')
		userCommands = strings.Trim(userCommands, "\n")
		commandLen := len(strings.Split(userCommands, " "))
		fmt.Fprintln(conn, userCommands)

		// Get Request
		//severReply, _ := bufio.NewReader(conn).ReadString('')
		severReply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(severReply)

		// User can stop program from client side
		if commandLen == 1 && strings.ToLower(userCommands) == "exit" {
			log.Println("Client Stopped")
			os.Exit(0)
		}
	}

}
