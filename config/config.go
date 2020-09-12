package main

import (
	"bufio"
	"fmt"
	"os"
)

type config struct {
	ID   int
	port string
	IP   string
	minD int
	maxD int
}

func ReadFile(ConfigFile string) []config {
	type minD int

	type maxD int

	//open and read config file
	ConfigFile, err := os.Open("config.txt") //assuming this will be file name
	if err != nil {
		fmt.Println(err)
	}
	defer ConfigFile.Close()

	//read the files contents
	scanner := bufio.NewScanner(ConfigFile)
	//parse out each word individually
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		//parse out details of the config file to then populate the necessary fields
		fmt.Println(scanner.Text())

	}
	//pseudocode
	/*
	   parse line 1 from file and populate minD and maxD
	   readfile by line: AFTER line 1 break each line into 3 strings: populate fields from this, and set minD and maxD

	*/
	return nil // finish writing return statement
}
