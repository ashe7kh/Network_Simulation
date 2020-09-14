package main

import (
	"../config"
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

//type cfile = config.Config
type cfile = config.Config
func readF(s string) []cfile{
	return config.ReadFile(s)
}
type message struct {
	Content string
	Time    string
}

func Messageprint(ID int, Message string) { //function that prints the message on the server side

	fmt.Println("\n -------------------------- \n --- New Incoming Message --- \n -------------------------- \n")
	fmt.Println("Message from:" + string(ID)) //unique identity of sender
	fmt.Println("Message Content: " + Message)
	t := time.Now()
	myTime := t.Format(time.RFC850) + "\n"
	fmt.Println("Confirmed received at: " + myTime)
	fmt.Println("\n -------------------------- \n")

}

//takes the message and then sends it to be printed
func UnicastReceive(config cfile, message string) {
	var identifier int
	identifier = config.ID
	Messageprint(identifier, message)
}

func main() {
	var config cfile
	config = readF("config.txt")[0] //two processes, received from process 1 => ID = 1 => index = 0
	//declare type of communication as well as the port of access
	address := config.IP + ":" + config.Port // need to differentiate between both ID's and IP's
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		//listen for communication
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		for {
			//read the incoming information with protocol in case of error
			netData, err := bufio.NewReader(conn).ReadString('\n')

			if err != nil {
				fmt.Println(err)
				return
			}

			//termination protocol allows the client to end connection manually
			if strings.TrimSpace(netData) != "END" {
				//if termination protocol is not called then receive and print the message
				UnicastReceive(config, netData)
			} else {
				fmt.Println("Exiting TCP server!")
				return
			}
		}
		return
	}
}
