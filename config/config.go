package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	MinD int
	MaxD int
	ID   int
	port string
	IP   string
}

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
	type file string
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
		configs[counter].port = temp[1]
		configs[counter].IP = temp[2]

		counter += 1
	}
	//pseudocode
	/*
	   parse line 1 from file and populate minD and maxD
	   readfile by line: AFTER line 1 break each line into 3 strings: populate fields from this, and set minD and maxD

	*/
	return configs // finish writing return statement
}
