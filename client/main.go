package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gorilla/http"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %v $URL", os.Args[0])
	}
	//if _, err := http.Get(os.Stdout, os.Args[1]); err != nil {
	//	log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	//}

	url := os.Args[1]
	log.Printf("%s\n", url)
	status, body, reader, err := http.DefaultClient.Get(url, nil)
	if err != nil {
		log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
	}

	log.Println(status.String())
	log.Printf("%s\n", body)

	if reader != nil {
		//defer reader.Close()
		buffer, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
		}
		log.Printf("%s\n", string(buffer))
	}
}
