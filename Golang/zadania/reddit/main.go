package main

import (
	"io"
	"log"
	"os"
	"reddit/fetcher"
	"strings"
	"sync"
)

func main() {
	var urls = []string{"https://www.reddit.com/r/golang.json",
		"http://www.reddit.com/r/Polska.json",
		"http://www.reddit.com/r/poland.json",
		"http://www.reddit.com/r/movies.json"}

	//Saving data to a file
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	//err = oneURL(file, urls[0])
	//if err != nil {
	//	log.Print(err)
	err = multipleURLS(file, urls)
	if err != nil {
		log.Print(err)
	}
}

func oneURL(file io.Writer, url string) error {
	f := fetcher.Response{}

	//Fetching data
	err := f.Fetch(url)
	if err != nil {
		return err
	}
	err = f.Save(file, url)
	if err != nil {
		return err
	}
	return nil
}

func multipleURLS(file io.Writer, urls []string) (err error) {
	var fetchers []fetcher.Response
	var str strings.Builder
	wg := &sync.WaitGroup{}
	//Creating fetcher array
	for range urls {
		fetchers = append(fetchers, fetcher.Response{})
	}
	for i, url := range urls {
		wg.Add(1)
		i := i
		go func(url string, wg *sync.WaitGroup) {
			err := fetchers[i].Fetch(url)
			if err != nil {
				log.Print(err)
			}
			str.WriteString(fetchers[i].ReadData(url))
			if err != nil {
				log.Print(err)
			}
			wg.Done()
		}(url, wg)
	}
	wg.Wait()

	_, err = file.Write([]byte(str.String()))
	if err != nil {
		return err
	}
	return nil
}
