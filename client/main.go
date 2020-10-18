package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %v $URL", os.Args[0])
	}

	u := os.Args[1]

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: jar,
	}

	err = login(client, u)
	if err != nil {
		log.Fatal(err)
	}

	err = secret(client, u)
	if err != nil {
		log.Fatal(err)
	}

	err = logout(client, u)
	if err != nil {
		log.Fatal(err)
	}

	err = secret(client, u)
	if err != nil {
		log.Fatal(err)
	}
}

func get(client *http.Client, u string) {
	resp, err := client.Get(u)
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

func login(client *http.Client, u string) error {
	log.Printf("login\n")
	u = u + "/login"

	resp, err := client.Get(u)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)

	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(string(body))

	ur, err := url.Parse(u)
	if err != nil {
		return err
	}
	cookies := client.Jar.Cookies(ur)

	cookieVal := ""
	for _, v := range cookies {
		fmt.Printf("%v\n", v)
		if v.Name == "cookie-name" {
			cookieVal = v.Value
		}
		fmt.Printf("%s\n", cookieVal)
	}

	return nil
}

func secret(client *http.Client, u string) error {
	log.Printf("secret\n")
	u = u + "/secret"

	resp, err := client.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}

func logout(client *http.Client, u string) error {
	log.Printf("logout\n")
	u = u + "/logout"

	resp, err := client.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return nil
}
