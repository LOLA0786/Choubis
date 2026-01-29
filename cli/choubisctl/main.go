package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: choubisctl reload|status")
		return
	}

	cmd := os.Args[1]

	switch cmd {

	case "reload":

		resp, err := http.Post(
			"http://localhost:9000/admin/reload",
			"application/json",
			nil,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Reload:", resp.Status)

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
