package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	uri, err := loadURI()
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("Enter command to execute on %s: ", uri)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(input) == "new uri" {
			fmt.Print("Enter new URI: ")
			uri, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			uri = strings.TrimSpace(uri)

			err = saveURI(uri)
			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		req, err := http.NewRequest("POST", uri, bytes.NewBufferString(input))
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

func loadURI() (string, error) {
	file, err := os.Open("config.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", fmt.Errorf("configuration file is empty")
}

func saveURI(uri string) error {
	file, err := os.Create("config.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(uri)
	if err != nil {
		return err
	}

	return nil
}
