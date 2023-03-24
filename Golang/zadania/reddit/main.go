package main

import (
	"log"
	"os"
	"reddit/fetcher"
)

func main() {
	const url string = "https://www.reddit.com/r/golang.json"
	f := fetcher.Response{}

	//Fetching data
	err := f.Fetch(url)
	if err != nil {
		log.Panic("Could not fetch from url: " + err.Error())
	}

	//Saving data to a file
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	err = f.Save(file)
	if err != nil {
		log.Print("Could not save file: " + err.Error())
	}
}
