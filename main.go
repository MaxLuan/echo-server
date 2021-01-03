package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"net"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	flag := os.Getenv("http")
	if flag == "true" {
		http.HandleFunc("/", httpHandler)
		log.Fatal(http.ListenAndServe(":3333", nil))
	} else {
		l, err := net.Listen(CONN_TYPE, CONN_HOST+ ":"+CONN_PORT)
		if err != nil {
			fmt.Println("Error listening: ", err.Error())
			os.Exit(1)
		}
		defer l.Close()
		fmt.Println("Listening on " + CONN_HOST + "：" + CONN_PORT)
		for {
			fmt.Println("about to accpet.")
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			fmt.Println("Accepted connection.")
			go handleRequest(conn)
		}
	}
}

func httpHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "echo from server: " + os.Getenv("ServerName"))
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println("received " + string(buf))
	conn.Write([]byte("echo from server: " + os.Getenv("ServerName")))
	conn.Close()
}
