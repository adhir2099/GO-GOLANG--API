package main

import (
	"bytes"
	"net/http"
)

func main() {

	url := "http://localhost:8098/user/update/1"
	var jsonStr = []byte(`{ "name": "User1", "email": "user1@user.com" }`)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
