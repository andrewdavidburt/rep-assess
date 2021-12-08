package packages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func get(url string) (response *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	
	if err != nil {
		return nil, err
	}

	return request(req);
}

func parseBody(resp *http.Response, model interface{}) (err error) {
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, model)
	return
}

func request(request *http.Request) (response *http.Response, err error) {
	client := http.Client{}
	
	fmt.Println(fmt.Sprintf("Requesting %s => %s", request.URL.RequestURI(), request.Method))
	
	return client.Do(request)
}
