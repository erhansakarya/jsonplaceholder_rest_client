package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client = resty.New()
)

func main() {
	response, err := client.SetTimeout(5*time.Second).R().SetHeader("Accept", "application/json").Get("https://jsonplaceholder.typicode.com/comments")
	if response.StatusCode() != http.StatusOK {
		fmt.Println(err.Error())
	}

	fmt.Println(string(response.Body()))
}
