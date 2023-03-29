package fetcher

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

func (f *Response) Save(writer io.Writer, url string) error {

	var str strings.Builder
	str.WriteString("================================================================\n")
	str.WriteString("Date created: " + time.Now().Format(time.DateTime) + "\n")
	str.WriteString("URL: " + url + "\n\n")

	//for _, i := range f.Data.Children {
	//	str.WriteString(i.Data.Title + "\n" + i.Data.URL + "\n\n")
	//}
	for id := 0; id < 5; id++ {
		str.WriteString(f.Data.Children[id].Data.Title + "\n" + f.Data.Children[id].Data.URL + "\n\n")
	}

	str.WriteString("================================================================\n")
	_, err := writer.Write([]byte(str.String()))
	if err != nil {
		return err
	}
	return nil
}
func (f *Response) ReadData(url string) string {
	var str strings.Builder
	str.WriteString("================================================================\n")
	str.WriteString("Date created: " + time.Now().Format(time.DateTime) + "\n")
	str.WriteString("URL: " + url + "\n\n")

	//for _, i := range f.Data.Children {
	//	str.WriteString(i.Data.Title + "\n" + i.Data.URL + "\n\n")
	//}
	for id := 0; id < 5; id++ {
		str.WriteString(f.Data.Children[id].Data.Title + "\n" + f.Data.Children[id].Data.URL + "\n\n")
	}

	str.WriteString("================================================================\n")
	return str.String()
}

func (f *Response) Fetch(url string) error {
	client := http.Client{Timeout: time.Second * 10}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &f)
	if err != nil {
		return err
	}
	return nil
}
