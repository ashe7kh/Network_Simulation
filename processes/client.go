package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
	"../config"

)

type cfile = config.Config

type message struct{
	Content string
	Time string
}

func printMessage(ID int, Message message){ //function that prints the email on the client side
	fmt.Println("\n -------------------------- \n ---Message Confirmed Sent--- \n -------------------------- \n")
	fmt.Println("Message Sent to Process 2: " )
	fmt.Println("Message Content: " + Message.Content)
	t := time.Now()
	timeSent := t.Format(time.RFC3339)
	timeSent = Message.Time
	fmt.Println("Confirmed sent at: " + Message.Time)
	fmt.Println("\n -------------------------- \n")
}
/*
func MessageToString(Message message) string{ //function that converts the email to a string
   var s string
   s += Message.Content + "\n"
   return s
}

/*
func MessageDelay(Config cfile, Message message) string {
   min := Config.MinD
   max := Config.MaxD
   //delay := make(chan int)

   //var d time.Duration = time.Duration(n * int(time.Millisecond))
   //tick := time.Tick(d)
   text := MessageToString(Message)

   return text


}

*/

func unicastSend(config cfile, m message) message{
	address := config.IP + ":" + config.port

	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, m.Content + "\n")

	printMessage(m)
	return m

}


func main(){
	//declare type of connection as well as the port from which it will occur

	for {
		var c cfile

		var m message
		timeChannel := make(chan int)
		//prompt user to construct email then read email
		//TODO add feature to detect incorrect format and provide instruction to correct
		//read what was written on the command line
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter your Message:")
		text, _ := reader.ReadString('\n')

		//termination protocol
		if strings.TrimSpace(text) == "END"{
			fmt.Println("Exiting TCP Client")
			return
		}

		//fill structure with message content
		m.Content = text

		//add a delay via message delay function
		min := 1//c.MinD
		max := 10//c.MaxD
		n := rand.Intn(max-min) + min
		timer := time.Duration(n)*time.Millisecond
		time.AfterFunc(time.Duration(timer),func(){
			unicastSend(c, m)
			})

	}
}
