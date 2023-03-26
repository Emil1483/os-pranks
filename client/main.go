package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command to execute: ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", "http://localhost:8080/execute", bytes.NewBufferString(command))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "text/plain")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Response:")
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		fmt.Println()
	}
}
