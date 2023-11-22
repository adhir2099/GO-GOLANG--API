package main

import (
	"bytes"
	"net/http"
)

func main() {

	url := "http://localhost:8098/user/new"
	var jsonStr = []byte(`{ "name": "newUser", "email": "userMail" }`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
