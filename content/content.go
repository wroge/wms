package content

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func From(url, user, password string) (b *bytes.Reader, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, password)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status Code Error %v", resp.StatusCode)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return bytes.NewReader(data), nil
}
