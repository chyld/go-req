package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Request sends HTTP requests
func Request(url string, ch chan map[string]interface{}) {
	client := http.Client{}
	resp, _ := client.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var data interface{}
	json.Unmarshal(body, &data)

	ch <- data.(map[string]interface{})
}
