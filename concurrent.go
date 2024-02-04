package main

import (
	"log"
	"net"
	"time"
)

func do(connection net.Conn) {
	buffer := make([]byte, 1024)
	
	//Read
	_, err := connection.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	//Processing: Sleep for a sec
	time.Sleep(10 * time.Second)

	//Write a response
	connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\nResponse: Success\r\n"))

	//Close the connection
	connection.Close()
}

//Multi Threaded TCP server
func main() {
	//Listen for TCP requests on the port 9999
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listener:", listener)

	for {
		log.Println("Awaiting connection from a client")
		//Accept the incoming connection
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Connection accepted: ", connection)
		
		//Spawns a new thread
		go do(connection)
	}
}