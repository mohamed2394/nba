package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ThisForThat struct {
	This string `json:"this"`
	That string `json:"that"`
}

func main() {
	var thisorthat ThisForThat
	data, err := http.Get("https://itsthisforthat.com/api.php?json")
	if err != nil {
		log.Fatal(err)
	}
	dcd := json.NewDecoder(data.Body).Decode(&thisorthat)
	if dcd != nil {
		log.Fatal(dcd)
	}
	fmt.Printf("so basically it's %s for %s", thisorthat.This, thisorthat.That)

}
