package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var centerpoint_url string = "http://127.0.0.1:5000/centerpoint"

func centerpoint(ps [][2]int) {
	json_data, err := json.Marshal(ps)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(centerpoint_url, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	_ = bodyString
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	ps := [][2]int{}

	for i := 0; i < 100; i++ {
		x := rand.Intn(200) - 100
		y := rand.Intn(200) - 100

		ps = append(ps, [2]int{x, y})
	}

	centerpoint(ps)
}
