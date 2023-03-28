package main

import (
	"io"
	"log"
	"os"
	"reddit/fetcher"
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

	//err := oneURL(urls[0])
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
	log.Print(f)
	err = f.Save(file, url)
	if err != nil {
		return err
	}
	return nil
}

func multipleURLS(file io.Writer, urls []string) (err error) {
	var fetchers []fetcher.Response
	//Creating fetcher array
	for range urls {
		fetchers = append(fetchers, fetcher.Response{})
	}
	//for i, url := range urls {
	//	err := fetchers[i].Fetch(url)
	//	if err != nil {
	//		return err
	//	}
	//	err = fetchers[i].Save(file, url)
	//	if err != nil {
	//		return err
	//	}
	//}
	for i := range urls {
		i := i
		go func() {
			err = oneURL(file, urls[i])
			if err != nil {
				log.Print(err)
			}
		}()
	}
	return nil
}
