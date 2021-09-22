package tools

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//DoGetRequest is a common function, it helps to avoid repeat code every time
//that we need request some endpoint.
func DoGetRequest(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		return err
	}

	return nil
}
