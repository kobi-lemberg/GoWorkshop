// Access httpbin API to get our IP
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Reply is httpbin.org/ip reply
type Reply struct {
				  //Origin string // By convention JSON origin will map to field Origin
	IP string `json:"origin"` // map origin in json document to this field
}

func main() {
	resp, err := http.DefaultClient.Get("http://httpbin.org/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var reply Reply
	// Decode will fail on invalid JSON
	if err := dec.Decode(&reply); err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.IP)
}