package botutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/main/types"
)

func GetImageContent(keyword string, API_KEY string) types.Photos {
	getURL := fmt.Sprintf("https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=%s&text=hollow-knight&format=json", API_KEY)
	if keyword != "" {
		getURL = fmt.Sprintf("https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=%s&text=%s&format=json", keyword, API_KEY)
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

	var c types.Photos
	errUnmarshal := json.Unmarshal(bytes, &c)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return c
}
