package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"os"
)

func handleQuery(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Input query parameter is required", http.StatusBadRequest)
		return
	}

	fmt.Println("Input: ", input, "\n")
	cmd := exec.Command("python", "query_db.py", input)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")  // avoid encoding error

	// debug python
	debug := false
	if (debug) {
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			http.Error(w, "Failed to run python script: " + err.Error() + "\n\n" + stderr.String(), http.StatusInternalServerError)
			return
		}
	    fmt.Fprintf(w, "No error")
		return
	}

	output, err := cmd.Output()
	fmt.Fprintf(w, "Output: %s", output)
	if err != nil {
		http.Error(w, "Error: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/query", handleQuery)
	fmt.Println("Server started on :8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
