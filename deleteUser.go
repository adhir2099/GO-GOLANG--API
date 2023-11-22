package main

import (
	"bytes"
	"net/http"
)

func main() {

	url := "http://localhost:8098/user/delete/1"
	var jsonStr = []byte(``)
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
