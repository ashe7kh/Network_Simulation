package main

import (
	"bufio"
	"fmt"
	"io"
	"../config"
	"math/rand"
	"net"
	"strings"
	"time"
)

type message struct{
	Content string
	Time string
}

//type cfile = config.Config
var c cfile

func printMessage(Message message){ //function that prints the email on the server side

	fmt.Println("\n -------------------------- \n --- New Incoming Message --- \n -------------------------- \n")
	fmt.Println("Message Content: " + Message.Content)
	fmt.Println("Confirmed sent at: " + Message.Time)
	fmt.Println("\n -------------------------- \n")
}

func MessageToString(Message message) string{ //function that converts the email to a string to be sent to the client

	var s string
	s = "\n -------------------------\n --- message Client Copy --- \n ------------------------- \n"
	s += "Content: " + Message.Content + "\n"
	t := time.Now()
	timeRecieved := t.Format(time.RFC3339)
	s += "Time recieved: " + timeRecieved + "\n -------------------------\n"
	return s
}

func MessageDelay(Config cfile, Message message) string{

	min := Config.MinD
	max := Config.MaxD
	n := rand.Intn(max-min) + min
	printMessage(Message)
	var d time.Duration = time.Duration(n * int(time.Millisecond))
	tick := time.Tick(d)

	//convert the email back into a string to be sent back to the client
	a := MessageToString(Message)
	return a
}

func parser(netdata string) message{
	var m message
	m.Content = netdata
	t := time.Now()
	myTime := t.Format(time.RFC3339) + "\n"
	m.Time = myTime
	return m

}

func main() {
	//declare type of communication as well as the port of access
	ln, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	//dont close until exited
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		//read info from message
		for {
			//read the incoming information with protocol in case of error
			netData, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			//parse out each portion of the message and then populate each field of message structure
			for{
				var messages []message
				input := strings.NewReader(netData)
				t := time.Now()
				myTime := t.Format(time.RFC3339) + "\n"
				messages[0].Content = string(input)

				MessageDelay(Config config, )
			}

			//print the email on the server

			io.WriteString(conn, a + "\n")



			//termination protocol allows the client to end connection manually
			if strings.TrimSpace(netData) == "END" {
				fmt.Println("Exiting TCP server!")
				return
			}
		}
	}
}
