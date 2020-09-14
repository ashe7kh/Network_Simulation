package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type message struct{
	Content string
	Time string
}

//todo allow for the creation of numerous unique ID's and IP's to foster multicast functionality
type Config struct {
	MinD int
	MaxD int
	ID   int
	Port string
	IP   string
}

//extract delay values and populate the structure for future reference
func getDelayVals(firstLine []string) [2]int { //to tidy up ReadFile fxn; for the first line with delay values
	if len(firstLine) != 2 {
		panic(1)
	}

	var Delays [2]int
	var err error
	Delays[0], err = strconv.Atoi(firstLine[0]) //max delay
	if err != nil {
		panic(err)
	}
	Delays[1], err = strconv.Atoi(firstLine[1]) //min delay
	if err != nil {
		panic(err)
	}
	return Delays
}

func ReadFile(FileName string) []Config {

	var configs []Config

	//open and read config file
	ConfigFile, err := os.Open(FileName) //assuming this will be file name
	if err != nil {
		fmt.Println(err)
	}
	defer ConfigFile.Close()

	//read the files contents
	scanner := bufio.NewScanner(ConfigFile)
	//parse out file line by line
	DelayVals := getDelayVals(strings.Split(scanner.Text(), " ")) //delay values, 0th index = min, 1st = max
	scanner.Scan()                                                //move out of the first line
	counter := 0
	for scanner.Scan() {
		//parse out details of the config file to then populate the necessary fields
		temp := strings.Split(scanner.Text(), " ")
		configs[counter].MinD = DelayVals[0]
		configs[counter].MaxD = DelayVals[1]
		configs[counter].ID, err = strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		configs[counter].Port = temp[1]
		configs[counter].IP = temp[2]

		counter += 1
	}
	return configs // finish writing return statement
}
