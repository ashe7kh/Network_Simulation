package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	//declare type of connection as well as the port from which it will occur
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {

		//prompt user to construct email then read email
		//TODO add feature to detect incorrect format and provide instruction to correct
		//read what was written on the command line then send it through connection previosly established
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter your email with format To; From; Title; Content: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		//receiving confirmation from server
		//the client then receives a copy of their email as confirmation of its delivery
		message, _ := bufio.NewReader(conn).ReadString('|')
		fmt.Print("\n" + message)

		//disconection protocol
		if len(message) >1  {
			fmt.Println("TCP client exiting...")
			return
		}



	}
}
