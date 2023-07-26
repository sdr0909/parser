package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Proxy struct {
	IP        string   `json:"ip"`
	Port      string   `json:"port"`
	Protocols []string `json:"protocols"`
}

func main() {
	file, err := os.Open("proxies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var proxies []Proxy

	json.Unmarshal(byteValue, &proxies)

	for _, proxy := range proxies {
		for _, protocol := range proxy.Protocols {
			fmt.Printf("%s://%s:%s\n", protocol, proxy.IP, proxy.Port)
		}
	}
}
