package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Connection Succe")

	for {

		userCommands, _ := reader.ReadString('\n')
		userCommands = strings.Trim(userCommands, "\n")
		commandLen := len(strings.Split(userCommands, " "))
		fmt.Fprintln(conn, userCommands)

		if commandLen == 1 && strings.ToLower(userCommands) == "q" {
			log.Println("Client Stopped")
			os.Exit(0)
		}
	}

	fmt.Fprintln(conn, "I dialed you.")
	//fmt.Println(string(bs))

}
