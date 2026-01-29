package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, _ := http.Post(
		"http://localhost:9000/admin/policies",
		"application/json",
		nil,
	)

	fmt.Println("Reloaded:", resp.Status)
}
