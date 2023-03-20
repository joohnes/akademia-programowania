package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type fetcher struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func (f *fetcher) Save(writer io.Writer) error {

	var str strings.Builder
	str.WriteString("Date created: " + time.Now().Format(time.DateTime) + "\n\n")

	for _, i := range f.Data.Children {
		str.WriteString(i.Data.Title + "\n" + i.Data.URL + "\n\n")
	}

	_, err := writer.Write([]byte(str.String()))
	if err != nil {
		return err
	}
	return nil
}

func (f *fetcher) Fetch(url string) error {
	client := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	defer res.Body.Close()

	body, readError := io.ReadAll(res.Body)
	if readError != nil {
		return errors.New(readError.Error())
	}
	jsonErr := json.Unmarshal(body, &f)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func main() {
	const url string = "https://www.reddit.com/r/golang.json"
	f := fetcher{}

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
