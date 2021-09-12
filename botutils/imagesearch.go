package botutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"example.com/main/types"
)

func GetImageContent(keyword string, API_KEY string) types.ImageContent {
	getURL := fmt.Sprintf("https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=%s&text=aesthetic-wallpaper&format=json", API_KEY)
	if keyword != "" {
		getURL = fmt.Sprintf("https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=%s&text=%s-wallpaper&format=json", API_KEY, keyword)
	}
	resp, err := http.Get(getURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	bodyString = strings.ReplaceAll(bodyString, "jsonFlickrApi", "")
	bodyString = strings.ReplaceAll(bodyString, "(", "")
	bodyString = strings.ReplaceAll(bodyString, ")", "")
	log.Print(bodyString)
	bytes := []byte(bodyString)
	var c types.ImageContent
	errUnmarshal := json.Unmarshal(bytes, &c)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	log.Println(c.Photos.Photo[0].ID)
	return c
}
