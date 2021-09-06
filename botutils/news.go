package botutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/main/types"
)

func GetContent(keyword string, API_KEY string) types.NewsContent {
	getURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=%s", API_KEY)
	if keyword != "" {
		getURL = fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", keyword, API_KEY)
	}
	resp, err := http.Get(getURL)
	if err != nil {
		log.Fatal(err)
	}

	bytes, readErr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readErr != nil {
		log.Fatal(readErr)
	}

	var c types.NewsContent
	errUnmarshal := json.Unmarshal(bytes, &c)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return c
}
