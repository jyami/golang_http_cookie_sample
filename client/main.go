package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %v $URL", os.Args[0])
	}

	url := os.Args[1]

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	get(client, url)
}

func get(client *http.Client, url string) {
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Printf("%v\n", resp.Header)
	fmt.Println(string(body))
}
