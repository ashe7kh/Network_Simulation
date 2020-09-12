package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

type email struct{
	To string
	From string
	Title string
	Content string
	Time string
}

func printEmail(Email email){ //function that prints the email on the server side

	fmt.Println("\n -------------------------- \n --- New Incoming Email --- \n -------------------------- \n")
	fmt.Println("To: " + Email.To)
	fmt.Println("From: " + Email.From)
	fmt.Println("Title: " + Email.Title)
	fmt.Println("Content: " + Email.Content)
	fmt.Println("Confirmed sent at: " + Email.Time)
	fmt.Println("\n -------------------------- \n")
}

func EmailToString(Email email) string{ //function that converts the email to a string to be sent to the client
	var s string
	s = "\n -------------------------\n --- Email Client Copy --- \n ------------------------- \n"
	s += "To: " + Email.To + "\n"
	s += "From: " + Email.From + "\n"
	s += "Title: " + Email.Title + "\n"
	s += "Content: " + Email.Content + "\n"
	s += "Time sent: " + Email.Time + "\n -------------------------\n"
	return s
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
			//create an array of structure type email
			const num = 2
			var emails [2]email

			//read the incoming information with protocol in case of error
			netData, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}


			//parse out each portion of the email and then populate each field of email structure
			for i := 0; i < num; i++ {

				input := strings.Split(netData, ";")

				t := time.Now()
				myTime := t.Format(time.RFC3339) + "\n"

				emails[i] = email{input[0], input[1], input[2], input[3], myTime}
			}

			//print the email on the server
			printEmail(emails[0])

			//convert the email back into a string to be sent back to the client
			a := EmailToString(emails[0])

			//termination protocol allows the client to end connection manually
			if strings.TrimSpace(string(netData)) == "END" {
				fmt.Println("Exiting TCP server!")
				return
			}

			// send confirmation message by printing duplicate of the message on client side
			//this confirmation email triggers termination protocol
			io.WriteString(conn, a + "END\n|")
		}
	}
}
