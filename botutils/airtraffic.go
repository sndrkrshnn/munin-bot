package botutils

import (
	"encoding/json"
	"fmt"
	"go/types"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAirTraffic(keyword string, API_KEY string) types.Response {
	getURL := fmt.Sprintf("http://api.aviationstack.com/v1/flights?access_key=a417da4a76015b6bf49c624d91ca08a6")

	resp, err := http.Get(getURL)
	if err != nil {
		log.Fatal(err)
	}

	bytes, readErr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readErr != nil {
		log.Fatal(readErr)
	}

	var c types.Response
	errUnmarshal := json.Unmarshal(bytes, &c)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return c
}
