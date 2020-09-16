package main

import (
	"../config"
	"../message"
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

type confile = config.Config
type Message = message.Message

//function that prints the email on the client side
//prived time that the message was sent
func printMessage(ID int, message Message) {
	fmt.Println("\n -------------------------- \n ---Message Confirmed Sent--- \n -------------------------- \n")
	fmt.Println("Message Sent to:" + string(ID)) //add the unique id from config file
	fmt.Println("Message Content: " + message.Content)
	t := time.Now()
	timeSent := t.Format(time.RFC850)
	fmt.Println("Confirmed sent at: " + timeSent)
	timeSent = message.Time
	fmt.Println("\n -------------------------- \n")
}

func UnicastSend(config confile, m Message) {

	//obtain desired identifiers as well as IP addresses from config file
	//establish desired connection from config parameters
	//creates TCP connection on the client side
	address := config.IP + ":" + config.Port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//termination protocol
	if strings.TrimSpace(m.Content) != "END" {
		//send and then print the message on client APP side
		//todo have time be updated dynamically with the sending of the message not immediately after
		fmt.Fprintf(conn, m.Content+"\n")
		printMessage(config.ID, m)

	} else { //end communication
		fmt.Println("Exiting TCP Client")
		return
	}

}

//will extract the message itself from command line and then add delay prior to sending
func MessageParse(config confile, text string) {
	var c = config
	var m Message

	for {

		//create a string array to then parse out the message from the ID and IP info
		input := strings.Split(text, " ")

		//we only care about the information after declarations
		MessageActual := input[2:]

		//convert the array to a simple string
		text := strings.Join(MessageActual, " ")

		//fill structure with message content
		m.Content = text

		//add a delay via timer and AfterFunc
		//we determine delay amount by recalling the information from config file
		//min := c.MinD
		//max := c.MaxD
		n := rand.Intn(config.MaxD-config.MinD) + config.MinD

		timer := time.Duration(n) * time.Millisecond

		time.AfterFunc(timer, func() {
			UnicastSend(c, m)
		})

	}

}

func main() {
	var c confile
	//var m message
	c = config.ReadFile("config.txt")[1] //two processes, send to process 2 => ID = 2 => index = 1

	//prompt user to construct message then read message
	//TODO add feature to detect incorrect format and provide instruction to correct
	//read what was written on the command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Message:")

	//convert text to a string
	text, _ := reader.ReadString('\n')

	//then send the text to a parser function to only send desired information
	//parser also implements a delay in the form of a timer
	MessageParse(c, text)

}
