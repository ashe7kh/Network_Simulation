package main

import (
	"../config"
	"../message"
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type confile = config.Config
type Message = message.Message

//function that prints the email on the client side
//prived time that the message was sent
func printMessage(ID int, message string) {
	fmt.Println("\n ---------------------------- \n ---Message Confirmed Sent--- \n --------------------------- \n")
	fmt.Println("Message Sent to process:" + strconv.Itoa(ID))//add the unique id from config file
	fmt.Println("Message Content: " + message)
	t := time.Now()
	timeSent := t.Format(time.RFC850)
	fmt.Println("Confirmed sent at: " + timeSent)
	//timeSent = message.Time
	fmt.Println("\n --------------------------- \n")
}

func UnicastSend(conn net.Conn, m Message) {
	//call delay to waste some time

	//actually send the message
	fmt.Fprintf(conn, m.Content+"\n")

	//print message client side with time stamp
	printMessage(m.Local_ID, m.Content)

}

func delay(c confile){
	max := c.MaxD
	min := c.MinD
	//add timer to elapse a duration
	//call
	n := rand.Intn(max - min) + min
	ticker := time.NewTicker(time.Duration(n) * time.Millisecond)
	<- ticker.C
	ticker.Stop()
}

//will extract the message itself from command line and then add delay prior to sending
func MessageParse(text string) string{

	//create a string array to then parse out the message from the ID and IP info
	input := strings.Split(text, " ")

	//extract text after declarations
	MessageActual := input[2:]

	//convert the array to a simple string
	text = strings.Join(MessageActual, " ")

	return text

}

func main() {
	var c confile
	var m Message
	//var m message
	c = (config.ReadFile("config.txt"))[1] //two processes, send to process 2 => ID = 2 => index = 1

	//prompt user to construct message then read message
	//TODO add feature to detect incorrect format and provide instruction to correct
	//read what was written on the command line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Message:")

	//convert text to a string
	text, _ := reader.ReadString('\n')

	//create TCP channel
	address := c.IP + ":" + c.Port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//extract only te message from the command line
	m.Content = MessageParse(text)
	m.Local_ID = c.ID
	if strings.TrimSpace(text) != "END" {
		//add delay before sending
		delay(c)

		//send the message
		UnicastSend(conn, m)
		m.Time = time.Now()

	} else { //end communication
		fmt.Println("Exiting TCP Client")
		return
	}

}
